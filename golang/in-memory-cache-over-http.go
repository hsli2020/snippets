package main  // https://github.com/healeycodes/in-memory-cache-over-http

import (
	api "healeycodes/in-memory-cache-over-http/api"
)

func main() {
	api.Listen()
}

package api

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"healeycodes/in-memory-cache-over-http/cache"
)

var s *cache.Store

// Listen on PORT. Defaults to 8000
func Listen() {
	new()
	setup()
	start()
}

// New store
func new() {
	size, _ := strconv.Atoi(getEnv("SIZE", "0"))
	s = cache.Service(size)
}

// Setup path handlers
func setup() {
	http.HandleFunc("/get", handle(Get))
	http.HandleFunc("/set", handle(Set))
	http.HandleFunc("/delete", handle(Delete))
	http.HandleFunc("/checkandset", handle(CheckAndSet))
	http.HandleFunc("/increment", handle(Increment))
	http.HandleFunc("/decrement", handle(Decrement))
	http.HandleFunc("/append", handle(Append))
	http.HandleFunc("/prepend", handle(Prepend))
	http.HandleFunc("/stats", handle(Stats))
	http.HandleFunc("/flush", handle(Flush))
}

// Start http
func start() {
	port := getEnv("PORT", ":8000")
	fmt.Println("Listening on", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		panic(err)
	}
}

// Get a key from the store
// Status code: 200 if present, else 404
// e.g. ?key=foo
func Get(w http.ResponseWriter, r *http.Request) {
	value, exist := s.Get(r.URL.Query().Get("key"))
	if !exist {
		http.Error(w, "", 404)
		return
	}
	w.Header().Set("content-type", "text/plain")
	w.Write([]byte(value))
}

// Set a key in the store
// Status code: 204
func Set(w http.ResponseWriter, r *http.Request) {
	s.Set(
		r.URL.Query().Get("key"),
		r.URL.Query().Get("value"),
		getExpire(r.URL.Query().Get("expire")))
	w.WriteHeader(http.StatusNoContent)
}

// Delete a key in the store
// Status code: 204
func Delete(w http.ResponseWriter, r *http.Request) {
	s.Delete(r.URL.Query().Get("key"))
	w.WriteHeader(http.StatusNoContent)
}

// CheckAndSet a key in the store if it matches the compare value
// Status code: 204 if matches else 400
func CheckAndSet(w http.ResponseWriter, r *http.Request) {
	if s.CheckAndSet(
		r.URL.Query().Get("key"),
		r.URL.Query().Get("value"),
		getExpire(r.URL.Query().Get("expire")),
		r.URL.Query().Get("compare")) == true {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

// Increment a key in the store by an amount. If key missing, set the amount
// Status code: 204 if incrementable else 400
func Increment(w http.ResponseWriter, r *http.Request) {
	if err := s.Increment(
		r.URL.Query().Get("key"),
		r.URL.Query().Get("value"),
		getExpire(r.URL.Query().Get("expire"))); err == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

// Decrement a key in the store by an amount. If key missing, set the amount
// Status code: 204 if decrementable else 400
func Decrement(w http.ResponseWriter, r *http.Request) {
	if err := s.Decrement(
		r.URL.Query().Get("key"),
		r.URL.Query().Get("value"),
		getExpire(r.URL.Query().Get("expire"))); err == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

// Append to a key in the store
// Status code: 204
func Append(w http.ResponseWriter, r *http.Request) {
	s.Append(
		r.URL.Query().Get("key"),
		r.URL.Query().Get("value"),
		getExpire(r.URL.Query().Get("expire")))
	w.WriteHeader(http.StatusNoContent)
}

// Prepend to a key in the store
// Status code: 204
func Prepend(w http.ResponseWriter, r *http.Request) {
	s.Prepend(
		r.URL.Query().Get("key"),
		r.URL.Query().Get("value"),
		getExpire(r.URL.Query().Get("expire")))
}

// Flush all keys
func Flush(w http.ResponseWriter, r *http.Request) {
	s.Flush()
	w.WriteHeader(http.StatusNoContent)
}

// Stats of the cache
// Status code: 200
func Stats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(s.Stats()))
}

// Middleware
func handle(f func(http.ResponseWriter, *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		s.Mutex.Lock()
		defer s.Mutex.Unlock()
		f(w, r)
	}
}

// Safely get the expire, 0 if error
func getExpire(attempt string) int {
	value, err := strconv.Atoi(attempt)
	if err != nil {
		return 0
	}
	return value
}

// Gets an ENV variable else returns fallback
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

package cache

import (
	"container/list"
	"fmt"
	"strconv"
	"sync"
	"time"
)

// Store contains an LRU Cache
type Store struct {
	Mutex *sync.Mutex
	store map[string]*list.Element
	ll    *list.List
	max   int // Zero for unlimited
}

// Node maps a value to a key
type Node struct {
	key    string
	value  string
	expire int // Unix time
}

// Service returns an empty store
func Service(max int) *Store {
	s := &Store{
		Mutex: &sync.Mutex{},
		store: make(map[string]*list.Element),
		ll:    list.New(),
		max:   max,
	}
	return s
}

// Set a key
func (s *Store) Set(key string, value string, expire int) {
	current, exist := s.store[key]
	if exist != true {
		s.store[key] = s.ll.PushFront(&Node{
			key:    key,
			value:  value,
			expire: expire,
		})
		if s.max != 0 && s.ll.Len() > s.max {
			s.Delete(s.ll.Remove(s.ll.Back()).(*Node).key)
		}
		return
	}
	current.Value.(*Node).value = value
	current.Value.(*Node).expire = expire
	s.ll.MoveToFront(current)
}

// Get a key
func (s *Store) Get(key string) (string, bool) {
	current, exist := s.store[key]
	if exist {
		expire := int64(current.Value.(*Node).expire)
		if expire == 0 || expire > time.Now().Unix() {
			s.ll.MoveToFront(current)
			return current.Value.(*Node).value, true
		}
	}
	return "", false
}

// Delete an item
func (s *Store) Delete(key string) {
	current, exist := s.store[key]
	if exist != true {
		return
	}
	s.ll.Remove(current)
	delete(s.store, key)
}

// CheckAndSet a key. Sets only if the compare matches. Set the key if it doesn't exist
func (s *Store) CheckAndSet(key string, value string, expire int, compare string) bool {
	current, exist := s.store[key]
	if !exist || current.Value.(*Node).value == compare {
		s.Set(key, value, expire)
		return true
	}
	return false
}

// Increment a key by an amount. Both value and amount should be integers.
// If doesn't exist, set to amount
func (s *Store) Increment(key string, value string, expire int) error {
	current, exist := s.store[key]
	if !exist {
		s.Set(key, value, expire)
	}

	y, err := strconv.Atoi(value)
	if err != nil {
		return err
	}
	x, err := strconv.Atoi(current.Value.(*Node).value)
	if err != nil {
		return err
	}
	s.Set(key, strconv.Itoa(x+y), expire)
	return nil
}

// Decrement a key by an amount. Both value and amount should be integers.
// If doesn't exist, set to amount
func (s *Store) Decrement(key string, value string, expire int) error {
	current, exist := s.store[key]
	if !exist {
		s.Set(key, value, expire)
	}

	y, err := strconv.Atoi(value)
	if err != nil {
		return err
	}
	x, err := strconv.Atoi(current.Value.(*Node).value)
	if err != nil {
		return err
	}
	s.Set(key, strconv.Itoa(x-y), expire)
	return nil
}

// Append to a key
func (s *Store) Append(key string, value string, expire int) {
	current, exist := s.store[key]
	if !exist {
		s.Set(key, value, expire)
		return
	}
	s.Set(key, current.Value.(*Node).value+value, expire)
}

// Prepend to a key
func (s *Store) Prepend(key string, value string, expire int) {
	current, exist := s.store[key]
	if !exist {
		s.Set(key, value, expire)
		return
	}
	s.Set(key, value+current.Value.(*Node).value, expire)
}

// Flush all keys
func (s *Store) Flush() {
	s.store = make(map[string]*list.Element)
	s.ll = list.New()
}

// Stats returns up-to-date information about the cache
func (s *Store) Stats() string {
	// TODO (healeycodes)
	// Use json package here
	return fmt.Sprintf(`{"keyCount": %v, "maxSize": %v}`, len(s.store), s.max)
}
