// Package rest allows changing the HTTP method via a query string.
package rest

import "net/http"
import "strings"

// Handler will update the HTTP request type.
func Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			// Clean the query string
			values := r.URL.Query()
			method := values.Get("_method")
			values.Del("_method")
			r.URL.RawQuery = values.Encode()

			// Set the method
			if len(method) > 0 {
				r.Method = strings.ToUpper(method)
			}
		}

		next.ServeHTTP(w, r)
	})
}

// Package logrequest provides an http.Handler that logs when a request is
// made to the application and lists the remote address, the HTTP method, and the URL.
package logrequest

import "fmt"
import "net/http"
import "time"

// Handler will log the HTTP requests.
func Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(time.Now().Format("2006-01-02 03:04:05 PM"), r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

// Package acl provides http.Handlers to prevent access to pages for
// authenticated users and for non-authenticated users.
package acl

import "net/http"
import "github.com/blue-jay/blueprint/lib/flight"

// DisallowAuth does not allow authenticated users to access the page.
func DisallowAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := flight.Context(w, r)

		// If user is authenticated, don't allow them to access the page
		if c.Sess.Values["id"] != nil {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		h.ServeHTTP(w, r)
	})
}

// DisallowAnon does not allow anonymous users to access the page.
func DisallowAnon(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := flight.Context(w, r)

		// If user is not authenticated, don't allow them to access the page
		if c.Sess.Values["id"] == nil {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		h.ServeHTTP(w, r)
	})
}

/* Source: https://github.com/heppu/simple-cors */
package cors

import "net/http"

const (
	options           string = "OPTIONS"
	allow_origin      string = "Access-Control-Allow-Origin"
	allow_methods     string = "Access-Control-Allow-Methods"
	allow_headers     string = "Access-Control-Allow-Headers"
	allow_credentials string = "Access-Control-Allow-Credentials"
	expose_headers    string = "Access-Control-Expose-Headers"
	credentials       string = "true"
	origin            string = "Origin"
	methods           string = "POST, GET, OPTIONS, PUT, DELETE, HEAD, PATCH"

	// If you want to expose some other headers add it here
	headers string = "Accept, Accept-Encoding, Authorization, Content-Length, Content-Type, X-CSRF-Token"
)

// Handler will allow cross-origin HTTP requests
func Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set allow origin to match origin of our request or fall back to *
		if o := r.Header.Get(origin); o != "" {
			w.Header().Set(allow_origin, o)
		} else {
			w.Header().Set(allow_origin, "*")
		}

		// Set other headers
		w.Header().Set(allow_headers, headers)
		w.Header().Set(allow_methods, methods)
		w.Header().Set(allow_credentials, credentials)
		w.Header().Set(expose_headers, headers)

		// If this was preflight options request let's write empty ok response and return
		if r.Method == options {
			w.WriteHeader(http.StatusOK)
			w.Write(nil)
			return
		}

		next.ServeHTTP(w, r)
	})
}

package pprofhandler

import "net/http"
import "net/http/pprof"
import "github.com/gorilla/context"
import "github.com/julienschmidt/httprouter"

// Handler routes the pprof pages using httprouter
func Handler(w http.ResponseWriter, r *http.Request) {
	p := context.Get(r, "params").(httprouter.Params)

	switch p.ByName("pprof") {
	case "/cmdline":
		pprof.Cmdline(w, r)
	case "/profile":
		pprof.Profile(w, r)
	case "/symbol":
		pprof.Symbol(w, r)
	default:
		pprof.Index(w, r)
	}
}

package acl

import "net/http"
import "app/shared/session"

// DisallowAuth does not allow authenticated users to access the page
func DisallowAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get session
		sess := session.Instance(r)

		// If user is authenticated, don't allow them to access the page
		if sess.Values["id"] != nil {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		h.ServeHTTP(w, r)
	})
}

// DisallowAnon does not allow anonymous users to access the page
func DisallowAnon(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get session
		sess := session.Instance(r)

		// If user is not authenticated, don't allow them to access the page
		if sess.Values["id"] == nil {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		h.ServeHTTP(w, r)
	})
}
