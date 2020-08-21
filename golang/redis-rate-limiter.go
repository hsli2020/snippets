// https://github.com/mauricio/redis-rate-limiter

// limiter.go
package redis_rate_limiter

import (
	"context"
	"time"
)

// Request defines a request that needs to be checked if it will be rate limited or not.
// The `Key` is the identifier you're using for the client making calls. This could be a 
// user/account ID if the user is signed into your application, the IP of the client making 
// requests (this might not be reliable if you're not behind a proxy like Cloudflare, as clients 
// can try to spoof IPs). The `Key` should be the same for multiple calls of the same client so 
// we can correctly identify that this is the same app calling anywhere.
// `Limit` is the amount of requests the client is allowed to make over the `Duration` period. 
// If you set this to 100 and `Duration` to `1m` you'd have at most 100 requests over a minute.
type Request struct {
	Key      string
	Limit    uint64
	Duration time.Duration
}

// State is the result of evaluating the rate limit, either `Deny` or `Allow` a request.
type State int64

const (
	Deny  State = 0
	Allow       = 1
)

// Result represents the response to a check if a client should be rate limited or not. 
// The `State` will be either `Allow` or `Deny`, `TotalRequests` holds the number of requests 
// this specific caller has already made over the current period of time and `ExpiresAt` 
// defines when the rate limit will expire/roll over for clients that have gone over the limit.
type Result struct {
	State         State
	TotalRequests uint64
	ExpiresAt     time.Time
}

// Strategy is the interface the rate limit implementations must implement to be used, 
// it takes a `Request` and returns a `Result` and an `error`, any errors the rate-limiter 
// finds should be bubbled up so the code can make a decision about what it wants to do 
// with the request.
type Strategy interface {
	Run(ctx context.Context, r *Request) (*Result, error)
}

// sorted_set_counter.go
package redis_rate_limiter

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"strconv"
	"time"
)

var (
	_ Strategy = &sortedSetCounter{}
)

const (
	sortedSetMax = "+inf"
	sortedSetMin = "-inf"
)

func NewSortedSetCounterStrategy(client *redis.Client, now func() time.Time) Strategy {
	return &sortedSetCounter{
		client: client,
		now:    now,
	}
}

type sortedSetCounter struct {
	client *redis.Client
	now    func() time.Time
}

// Run this implementation uses a sorted set that holds an UUID for every request with a score 
// that is the time the request has happened. This allows us to delete events from *before* the 
// current window, if the window is 5 minutes, we want to remove all events from before 5 minutes 
// ago, this way we make sure we roll old requests off the counters creating a rolling window 
// for the rate limiter. So, if your window is 100 requests in 5 minutes and you spread the 
// load evenly across the minutes, once you hit 6 minutes the requests you made on the first 
// minute have now expired but the other 4 minutes of requests are still valid.
// A rolling window counter is usually never 0 if traffic is consistent so it is very effective 
// at preventing bursts of traffic as the counter won't ever expire.
func (s *sortedSetCounter) Run(ctx context.Context, r *Request) (*Result, error) {
	now := s.now()
	expiresAt := now.Add(r.Duration)
	minimum := now.Add(-r.Duration)

	// first count how many requests over the period we're tracking on this rolling window 
    // so check wether we're already over the limit or not. this prevents new requests from 
    // being added if a client is already rate limited, not allowing it to add an infinite 
    // amount of requests to the system overloading redis.
	// if the client continues to send requests it also means that the memory for this specific 
    // key will not be reclaimed (as we're not writing data here) so make sure there is an 
    // eviction policy that will clear up the memory if the redis starts to get close to its 
    // memory limit.
	result, err := s.client.ZCount(ctx, r.Key, strconv.FormatInt(minimum.UnixMilli(), 10), sortedSetMax).Uint64()
	if err == nil && result >= r.Limit {
		return &Result{
			State:         Deny,
			TotalRequests: result,
			ExpiresAt:     expiresAt,
		}, nil
	}

	// every request needs an UUID
	item := uuid.New()

	p := s.client.Pipeline()

	// we then remove all requests that have already expired on this set
	removeByScore := p.ZRemRangeByScore(ctx, r.Key, "0", strconv.FormatInt(minimum.UnixMilli(), 10))

	// we add the current request
	add := p.ZAdd(ctx, r.Key, &redis.Z{
		Score:  float64(now.UnixMilli()),
		Member: item.String(),
	})

	// count how many non-expired requests we have on the sorted set
	count := p.ZCount(ctx, r.Key, sortedSetMin, sortedSetMax)

	if _, err := p.Exec(ctx); err != nil {
		return nil, errors.Wrapf(err, "failed to execute sorted set pipeline for key: %v", r.Key)
	}

	if err := removeByScore.Err(); err != nil {
		return nil, errors.Wrapf(err, "failed to remove items from key %v", r.Key)
	}

	if err := add.Err(); err != nil {
		return nil, errors.Wrapf(err, "failed to add item to key %v", r.Key)
	}

	totalRequests, err := count.Result()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to count items for key %v", r.Key)
	}

	requests := uint64(totalRequests)

	if requests > r.Limit {
		return &Result{
			State:         Deny,
			TotalRequests: requests,
			ExpiresAt:     expiresAt,
		}, nil
	}

	return &Result{
		State:         Allow,
		TotalRequests: requests,
		ExpiresAt:     expiresAt,
	}, nil
}

// http.go
package redis_rate_limiter

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	_            http.Handler = &httpRateLimiterHandler{}
	_            Extractor    = &httpHeaderExtractor{}
	stateStrings              = map[State]string{
		Allow: "Allow",
		Deny:  "Deny",
	}
)

const (
	rateLimitingTotalRequests = "Rate-Limiting-Total-Requests"
	rateLimitingState         = "Rate-Limiting-State"
	rateLimitingExpiresAt     = "Rate-Limiting-Expires-At"
)

// Extractor represents the way we will extract a key from an HTTP request, this could be
// a value from a header, request path, method used, user authentication information, any 
// information that is available at the HTTP request that wouldn't cause side effects if it 
// was collected (this object shouldn't read the body of the request).
type Extractor interface {
	Extract(r *http.Request) (string, error)
}

type httpHeaderExtractor struct {
	headers []string
}

// Extract extracts a collection of http headers and joins them to build the key that will be 
// used for rate limiting. You should use headers that are guaranteed to be unique for a client.
func (h *httpHeaderExtractor) Extract(r *http.Request) (string, error) {
	values := make([]string, 0, len(h.headers))

	for _, key := range h.headers {
		// if we can't find a value for the headers, give up and return an error.
		if value := strings.TrimSpace(r.Header.Get(key)); value == "" {
			return "", fmt.Errorf("the header %v must have a value set", key)
		} else {
			values = append(values, value)
		}
	}

	return strings.Join(values, "-"), nil
}

// NewHTTPHeadersExtractor creates a new HTTP header extractor
func NewHTTPHeadersExtractor(headers ...string) Extractor {
	return &httpHeaderExtractor{headers: headers}
}

// RateLimiterConfig holds the basic config we need to create a middleware http.Handler object 
// that performs rate limiting before offloading the request to an actual handler.
type RateLimiterConfig struct {
	Extractor   Extractor
	Strategy    Strategy
	Expiration  time.Duration
	MaxRequests uint64
}

// NewHTTPRateLimiterHandler wraps an existing http.Handler object performing rate limiting before
// sending the request to the wrapped handler. If any errors happen while trying to rate limit a 
// request or if the request is denied, the rate limiting handler will send a response to the 
// client and will not call the wrapped handler.
func NewHTTPRateLimiterHandler(originalHandler http.Handler, config *RateLimiterConfig) http.Handler {
	return &httpRateLimiterHandler{
		handler: originalHandler,
		config:  config,
	}
}

type httpRateLimiterHandler struct {
	handler http.Handler
	config  *RateLimiterConfig
}

func (h *httpRateLimiterHandler) writeRespone(writer http.ResponseWriter, status int, msg string, args ...interface{}) {
	writer.Header().Set("Content-Type", "text/plain")
	writer.WriteHeader(status)
	if _, err := writer.Write([]byte(fmt.Sprintf(msg, args...))); err != nil {
		fmt.Printf("failed to write body to HTTP request: %v", err)
	}
}

// ServeHTTP performs rate limiting with the configuration it was provided and if there were 
// not errors and the request was allowed it is sent to the wrapped handler. It also adds 
// rate limiting headers that will be sent to the client to make it aware of what state it 
// is in terms of rate limiting.
func (h *httpRateLimiterHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	key, err := h.config.Extractor.Extract(request)
	if err != nil {
		h.writeRespone(writer, http.StatusBadRequest, 
            "failed to collect rate limiting key from request: %v", err)
		return
	}

	result, err := h.config.Strategy.Run(request.Context(), &Request{
		Key:      key,
		Limit:    h.config.MaxRequests,
		Duration: h.config.Expiration,
	})

	if err != nil {
		h.writeRespone(writer, http.StatusInternalServerError, 
            "failed to run rate limiting for request: %v", err)
		return
	}

	// set the rate limiting headers both on allow or deny results so the client knows what is going on
	writer.Header().Set(rateLimitingTotalRequests, strconv.FormatUint(result.TotalRequests, 10))
	writer.Header().Set(rateLimitingState, stateStrings[result.State])
	writer.Header().Set(rateLimitingExpiresAt, result.ExpiresAt.Format(time.RFC3339))

	// when the state is Deny, just return a 429 response to the client and stop the request handling flow
	if result.State == Deny {
		h.writeRespone(writer, http.StatusTooManyRequests, 
            "you have sent too many requests to this service, slow down please")
		return
	}

	// if the request was not denied we assume it was allowed and call the wrapped handler.
	// by leaving this to the end we make sure the wrapped handler is only called once and 
    // doesn't have to worry about any rate limiting at all (it doesn't even have to know 
    // there was rate limiting happening for this request) as we have already set the headers, 
    // so when the handler flushes the response the headers above will be sent.
	h.handler.ServeHTTP(writer, request)
}

// counter.go
package redis_rate_limiter

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"time"
)

var (
	_ Strategy = &counterStrategy{}
)

const (
	keyThatDoesNotExist = -2
	keyWithoutExpire    = -1
)

func NewCounterStrategy(client *redis.Client, now func() time.Time) *counterStrategy {
	return &counterStrategy{
		client: client,
		now:    now,
	}
}

type counterStrategy struct {
	client *redis.Client
	now    func() time.Time
}

// Run this implementation uses a simple counter with an expiration set to the rate limit duration.
// This implementation is funtional but not very effective if you have to deal with bursty traffic as
// it will still allow a client to burn through it's full limit quickly once the key expires.
func (c *counterStrategy) Run(ctx context.Context, r *Request) (*Result, error) {

	// a pipeline in redis is a way to send multiple commands that will all be run together.
	// this is not a transaction and there are many ways in which these commands could fail
	// (only the first, only the second) so we have to make sure all errors are handled, this
	// is a network performance optimization.

	// here we try to get the current value and also try to set an expiration on it
	getPipeline := c.client.Pipeline()
	getResult := getPipeline.Get(ctx, r.Key)
	ttlResult := getPipeline.TTL(ctx, r.Key)

	if _, err := getPipeline.Exec(ctx); err != nil && !errors.Is(err, redis.Nil) {
		return nil, errors.Wrapf(err, "failed to execute pipeline with get and ttl to key %v", r.Key)
	}

	var ttlDuration time.Duration

	// we want to make sure there is always an expiration set on the key, so on every
	// increment we check again to make sure it has a TTl and if it doesn't we add one.
	// a duration of -1 means that the key has no expiration so we need to make sure there
	// is one set, this should, most of the time, happen when we increment for the
	// first time but there could be cases where we fail at the previous commands so we should
	// check for the TTL on every request.
	// a duration of -2 means that the key does not exist, given we're already here we should 
    // set an expiration to it anyway as it means this is a new key that will be incremented below.
	if d, err := ttlResult.Result(); err != nil || d == keyWithoutExpire || d == keyThatDoesNotExist {
		ttlDuration = r.Duration
		if err := c.client.Expire(ctx, r.Key, r.Duration).Err(); err != nil {
			return nil, errors.Wrapf(err, "failed to set an expiration to key %v", r.Key)
		}
	} else {
		ttlDuration = d
	}

	expiresAt := c.now().Add(ttlDuration)

	if total, err := getResult.Uint64(); err != nil && errors.Is(err, redis.Nil) {

	} else if total >= r.Limit {
		return &Result{
			State:         Deny,
			TotalRequests: total,
			ExpiresAt:     expiresAt,
		}, nil
	}

	incrResult := c.client.Incr(ctx, r.Key)

	totalRequests, err := incrResult.Uint64()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to increment key %v", r.Key)
	}

	if totalRequests > r.Limit {
		return &Result{
			State:         Deny,
			TotalRequests: totalRequests,
			ExpiresAt:     expiresAt,
		}, nil
	}

	return &Result{
		State:         Allow,
		TotalRequests: totalRequests,
		ExpiresAt:     expiresAt,
	}, nil
}
