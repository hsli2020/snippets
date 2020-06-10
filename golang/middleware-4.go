// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package middleware implements a simple middleware pattern for http handlers,
// along with implementations for some common middlewares.
package middleware

import "net/http"

// A Middleware is a func that wraps an http.Handler.
type Middleware func(http.Handler) http.Handler

// Chain creates a new Middleware that applies a sequence of Middlewares, so
// that they execute in the given order when handling an http request.
//
// In other words, Chain(m1, m2)(handler) = m1(m2(handler))
//
// A similar pattern is used in e.g. github.com/justinas/alice:
// https://github.com/justinas/alice/blob/ce87934/chain.go#L45
func Chain(middlewares ...Middleware) Middleware {
	return func(h http.Handler) http.Handler {
		for i := range middlewares {
			h = middlewares[len(middlewares)-1-i](h)
		}
		return h
	}
}

// Identity is a middleware that does nothing. It can be used as a helper when
// building middleware chains.
func Identity() Middleware {
	return func(h http.Handler) http.Handler {
		return h
	}
}

// nonce.go
package middleware

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

const (
	nonceCtxKey = "CSPNonce"
	nonceLen    = 20
)

func generateNonce() (string, error) {
	nonceBytes := make([]byte, nonceLen)
	if _, err := io.ReadAtLeast(rand.Reader, nonceBytes, nonceLen); err != nil {
		return "", fmt.Errorf("io.ReadAtLeast(rand.Reader, nonceBytes, %d): %v", nonceLen, err)
	}
	return base64.StdEncoding.EncodeToString(nonceBytes), nil
}

// timeout.go
package middleware

import (
	"context"
	"net/http"
	"time"
)

// Timeout returns a new Middleware that times out each request after the given
// duration.
func Timeout(d time.Duration) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, cancel := context.WithTimeout(r.Context(), d)
			defer cancel()
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// panic.go
package middleware

import (
	"net/http"

	"golang.org/x/pkgsite/internal/log"
)

// Panic returns a middleware that executes panicHandler on any panic
// originating from the delegate handler.
func Panic(panicHandler http.Handler) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if e := recover(); e != nil {
					log.Errorf(r.Context(), "middleware.Panic: %v", e)
					panicHandler.ServeHTTP(w, r)
				}
			}()
			h.ServeHTTP(w, r)
		})
	}
}

// errorreporting.go
package middleware

import (
	"fmt"
	"net/http"

	"cloud.google.com/go/errorreporting"
)

// ErrorReporting returns a middleware that reports any server errors using the
// report func.
func ErrorReporting(report func(errorreporting.Entry)) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w2 := &responseWriter{ResponseWriter: w}
			h.ServeHTTP(w2, r)
			if w2.status >= 500 {
				e := errorreporting.Entry{
					Error: fmt.Errorf("handler for %q returned status code %d", r.URL.Path, w2.status),
					Req:   r,
				}
				report(e)
			}
		})
	}
}

// accept_methods.go
package middleware

import "net/http"

// AcceptMethods serves 405 (Method Not Allowed) for any method not on the given list.
func AcceptMethods(methods ...string) Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			for _, m := range methods {
				if r.Method == m {
					h.ServeHTTP(w, r)
					return
				}
			}
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		})
	}
}
