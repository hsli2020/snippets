package main

import (
	"errors"
	"github.com/garyburd/redigo/redis"
	"net/http"
	"time"
)

type RateStore struct {
	Pool   *redis.Pool
	Limit  int64
	Header string
}

// NewRateStore returns a new RateStore.
// Depending on your setup or reverse proxy, you will need to set Header to
// inspect either "REMOTE_ADDR" or "X-Forwarded-For".
// Example:
//   	limitStore = NewRateStore(10, 1, "REMOTE_ADDR", "tcp", ":6380", "password")
//
// Note: You should spin up a second Redis instance if you already have a primary for other tasks.
func NewRateStore(idle int, limit int64, header, net, port, password string) *RateStore {
	return &RateStore{
		Pool: &redis.Pool{
			MaxIdle:     idle,
			IdleTimeout: 240 * time.Second,
			Dial: func() (c redis.Conn, err error) {
				c, err = redis.Dial(net, port)
				if err != nil {
					return nil, err
				}
				if password != "" {
					if _, err := c.Do("AUTH", password); err != nil {
						c.Close()
						return nil, err
					}
				}
				return c, err
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				_, err := c.Do("PING")
				return err
			},
		},
		Limit:  limit,
		Header: header,
	}
}

// RateLimit provides HTTP request limiting middleware. Requests are limited to Limit per second per IP.
// Requests that exceed the limit are served with HTTP 429 (Too Many Requests).
func (s *RateStore) RateLimit(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		rconn := s.Pool.Get()
		defer rconn.Close()

		path := r.URL.Path
		remoteIP := r.Header.Get(s.Header)
		// Invoke the next handler if the remote address is not set
		// (we cannot determine the rate without it)
		if remoteIP == "" {
			h.ServeHTTP(w, r)
			return
		}

		// INCR will increment an existing key (if any) else it creates a new one (at 1)
		current, err := rconn.Do("INCR", path+":"+remoteIP)
		if err != nil {
			serverError(w, r, err, 500)
			return
		}

		// Set a 1s expiry on newly instantiated counters
		if current.(int64) == 1 {
			_, err := rconn.Do("EXPIRE", path+":"+remoteIP, 1)
			if err != nil {
				serverError(w, r, err, 500)
				return
			}
		} else if current.(int64) > s.Limit {
			// Check if the returned counter exceeds the limit
			serverError(w, r, errors.New("Rate exceeded."), 429)
			return
		}

		// Invoke the next handler if we haven't hit the limit
		h.ServeHTTP(w, r)
		return
	})
}

// Close closes the current connection
func (s *RateStore) Close() {
	s.Pool.Close()
}
