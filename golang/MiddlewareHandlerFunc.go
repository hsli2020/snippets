/* Copyright (c) 2019 DoorDash */
package middleware

import "net/http"

// MiddlewareInterceptor intercepts an HTTP handler invocation, it is passed both response writer
// and request which after interception can be passed onto the handler function.

type MiddlewareInterceptor func(http.ResponseWriter, *http.Request, http.HandlerFunc)

// MiddlewareHandlerFunc builds on top of http.HandlerFunc, and exposes API to intercept
// with MiddlewareInterceptor.
//
// This allows building complex long chains without complicated struct manipulation

type MiddlewareHandlerFunc http.HandlerFunc

// Intercept returns back a continuation that will call install middleware to intercept
// the continuation call.

func (cont MiddlewareHandlerFunc) Intercept(mw MiddlewareInterceptor) MiddlewareHandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		mw(writer, request, http.HandlerFunc(cont))
	}
}

// MiddlewareChain is a collection of interceptors that will be invoked in there index order

type MiddlewareChain []MiddlewareInterceptor

// Handler allows hooking multiple middleware in single call.

func (chain MiddlewareChain) Handler(handler http.HandlerFunc) http.Handler {
	curr := MiddlewareHandlerFunc(handler)
	for i := len(chain) - 1; i >= 0; i-- {
		mw := chain[i]
		curr = curr.Intercept(mw)
	}

	return http.HandlerFunc(curr)
}

/* How to use
func HomeRouter(w http.ResponseWriter, r *http.Request) {
	// Handle your request
}

func NewElapsedTimeInterceptor() MiddlewareInterceptor {
    return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
        startTime := time.Now()
        defer func() {
            endTime := time.Now()
            elapsed := endTime.Sub(startTime)
            // Log the consumed time
        }()

        next(w, r)
    }
}

func NewRequestIdInterceptor() MiddlewareInterceptor {
    return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
        if r.Headers.Get("X-Request-Id") == "" {
            r.Headers.Set("X-Request-Id", generateRequestId())
        }

        next(w, r)
    }
}

// Some where when installing Hanlder
chain := MiddlewareHandlerFunc(HomeRouter).
  Intercept(NewElapsedTimeInterceptor()).
  Intercept(NewRequestIdInterceptor())


// Install it like regular HttpHandler
mux.Path("/home").HandlerFunc(http.HandlerFunc(chain))
*/
