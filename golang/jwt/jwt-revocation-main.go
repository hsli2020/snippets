package main  // https://github.com/adelowo/jwt-revocation

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

const (
	port = 8999
)

func main() {

	signingSecret := os.Getenv("JWT_SIGNING_SECRET")
	if len(signingSecret) == 0 {
		log.Fatal("JWT_SIGNING_SECRET not found in environment")
	}

	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}

	redis, err := NewRediClient(dsn)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	s := &store{
		RWMutex: sync.RWMutex{},
		data:    make(map[string]User),
	}

	mux.HandleFunc("/login", login(s, signingSecret))

	auth := requireAuth(s, redis, signingSecret)

	mux.HandleFunc("/user/logout", auth(logoutHandler(redis)))

	mux.HandleFunc("/user/profile", auth(userProfile))

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux); err != nil {
		log.Fatal(err)
	}
}

type apiGenericResponse struct {
	Message   string `json:"message"`
	Status    bool   `json:"status"`
	Timestamp int64  `json:"timestamp"`
}

func encode(w io.Writer, v interface{}) {
	_ = json.NewEncoder(w).Encode(v)
}

func login(s *store, signingSecret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var u User

		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encode(w, apiGenericResponse{
				Message:   "Invalid request body ",
				Status:    false,
				Timestamp: time.Now().Unix(),
			})
			return
		}

		if u.FullName == "" {
			w.WriteHeader(http.StatusBadRequest)
			encode(w, apiGenericResponse{
				Message:   "Please provide your name",
				Status:    false,
				Timestamp: time.Now().Unix(),
			})
			return
		}

		if u.Email == "" {
			w.WriteHeader(http.StatusBadRequest)
			encode(w, apiGenericResponse{
				Message:   "Please provide your email",
				Status:    false,
				Timestamp: time.Now().Unix(),
			})
			return
		}

		// no errors
		_ = s.Save(u)

		token, err := GenerateJWT(signingSecret, u)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			encode(w, apiGenericResponse{
				Message:   "Could not generate JWT",
				Status:    false,
				Timestamp: time.Now().Unix(),
			})
			return
		}

		w.Header().Set("X-JWT-APP", token)
		encode(w, apiGenericResponse{
			Message:   "You have been logged in successfully",
			Status:    true,
			Timestamp: time.Now().Unix(),
		})
	}
}

func userProfile(w http.ResponseWriter, r *http.Request) {
	encode(w, r.Context().Value(userContextID))
}

package main  // redis.go

import (
	"errors"

	"github.com/go-redis/redis/v7"
)

type Client struct {
	redis *redis.Client
}

const (
	blackListKey string = "jwt_blacklist"
)

var (
	errJTIBlacklisted = errors.New("jwt token has been blacklisted")
)

func NewRediClient(dsn string) (*Client, error) {

	client := redis.NewClient(&redis.Options{
		Addr:     dsn,
		Password: "",
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return &Client{redis: client}, nil
}

func (c *Client) IsBlacklisted(jti string) error {
	m, err := c.redis.SMembersMap(blackListKey).Result()
	if err != nil {
		return err
	}

	if _, ok := m[jti]; ok {
		return errJTIBlacklisted
	}

	return nil
}

func (c *Client) AddToBlacklist(jti string) error {
	_, err := c.redis.SAdd(blackListKey, jti).Result()
	return err
}

package main  // user.go

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type contextKey string

const (
	jtiContextID  = "jtiContextID"
	userContextID = "userContext"
)

type User struct {
	Email    string `json:"email"`
	FullName string `json:"full_name"`
}

type store struct {
	sync.RWMutex

	data map[string]User
}

func (s *store) Get(email string) (User, error) {
	s.RLock()
	defer s.RUnlock()

	user, ok := s.data[email]
	if !ok {
		return User{}, errors.New("user not found")
	}

	return user, nil
}

func (s *store) Save(u User) error {
	s.RLock()

	_, ok := s.data[u.Email]
	s.RUnlock()
	if ok {
		return nil
	}

	s.Lock()
	s.data[u.Email] = u
	s.Unlock()
	return nil
}

func GenerateJWT(signingSecret string, u User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nbf":   time.Now().Add(-1 * time.Second),
		"jti":   uuid.New().String(),
		"email": u.Email,
		"iss":   "JWT-revocation-app",
		"exp":   time.Now().Add(time.Hour * 168),
	})

	return token.SignedString([]byte(signingSecret))
}

func getToken(r *http.Request) string {
	return strings.Trim(strings.TrimLeft(r.Header.Get("Authorization"), "Bearer"), " ")
}

func requireAuth(store *store, redis *Client, signingSecret string) func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {

			token, err := jwt.Parse(getToken(r), func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}

				return []byte(signingSecret), nil
			})

			if err != nil {
				fmt.Println(err)
				w.WriteHeader(http.StatusBadRequest)
				encode(w, apiGenericResponse{
					Message:   "Invalid token provided",
					Status:    false,
					Timestamp: time.Now().Unix(),
				})
				return
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				jti := claims["jti"].(string)

				if err := redis.IsBlacklisted(jti); err != nil {

					w.WriteHeader(http.StatusUnauthorized)
					encode(w, apiGenericResponse{
						Message:   "Authorization denied",
						Status:    false,
						Timestamp: time.Now().Unix(),
					})
					return
				}

				user, err := store.Get(claims["email"].(string))
				if err != nil {
					encode(w, apiGenericResponse{
						Message:   "Could not complete request",
						Status:    false,
						Timestamp: time.Now().Unix(),
					})
					return
				}

				ctx := context.WithValue(r.Context(), userContextID, user)
				ctx = context.WithValue(ctx, jtiContextID, jti)

				next(w, r.WithContext(ctx))
				return
			}

			w.WriteHeader(http.StatusUnauthorized)
			encode(w, apiGenericResponse{
				Message:   "Token is invalid",
				Status:    false,
				Timestamp: time.Now().Unix(),
			})
		}
	}
}

func logoutHandler(redis *Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		jti := r.Context().Value(jtiContextID).(string)

		if err := redis.AddToBlacklist(jti); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			encode(w,apiGenericResponse{
				Message:   "Could not log you out",
				Status:    false,
				Timestamp: time.Now().Unix(),
			})
			return
		}

		encode(w,apiGenericResponse{
			Message:   "You have been successfully logged out",
			Status:    true,
			Timestamp: time.Now().Unix(),
		})
	}
}