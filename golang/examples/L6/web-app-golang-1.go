Serving Static Sites with Go
============================

File: static/example.html
-------------------------

<!doctype html>
<html>
<head>
  <meta charset="utf-8">
  <title>A static page</title>
  <link rel="stylesheet" href="/stylesheets/main.css">
</head>
<body>
  <h1>Hello from a static page</h1>
</body>
</html>

File: static/stylesheets/main.css
---------------------------------

body {color: #c0392b}

File: app.go
------------

package main

import (
  "log"
  "net/http"
)

func main() {
  fs := http.FileServer(http.Dir("static"))
  http.Handle("/", fs)

  log.Println("Listening...")
  http.ListenAndServe(":3000", nil)
}

Almost-Static Sites
===================

File: templates/layout.html
---------------------------

{{define "layout"}}
<!doctype html>
<html>
<head>
  <meta charset="utf-8">
  <title>{{template "title"}}</title>
  <link rel="stylesheet" href="/static/stylesheets/main.css">
</head>
<body>
  {{template "body"}}
</body>
</html>
{{end}}

File: templates/example.html
----------------------------

{{define "title"}}A templated page{{end}}

{{define "body"}}
<h1>Hello from a templated page</h1>
{{end}}

File: app.go
------------

package main

import (
  "html/template"
  "log"
  "net/http"
  "path"
)

func main() {
  fs := http.FileServer(http.Dir("static"))
  http.Handle("/static/", http.StripPrefix("/static/", fs))

  http.HandleFunc("/", serveTemplate)

  log.Println("Listening...")
  http.ListenAndServe(":3000", nil)
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
  lp := path.Join("templates", "layout.html")
  fp := path.Join("templates", r.URL.Path)

  tmpl, _ := template.ParseFiles(lp, fp)
  tmpl.ExecuteTemplate(w, "layout", nil)
}

Lastly, let us make the code a bit more robust. We should:

 - Send a 404 response if the requested template doesnt exist.
 - Send a 404 response if the requested template path is a directory.
 - Send a 500 response if the template.ParseFiles or template.ExecuteTemplate functions
        throw an error, and log the detailed error message.

File: app.go

package main

import (
  "html/template"
  "log"
  "net/http"
  "os"
  "path"
)

func main() {
  fs := http.FileServer(http.Dir("static"))
  http.Handle("/static/", http.StripPrefix("/static/", fs))
  http.HandleFunc("/", serveTemplate)

  log.Println("Listening...")
  http.ListenAndServe(":3000", nil)
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
  lp := path.Join("templates", "layout.html")
  fp := path.Join("templates", r.URL.Path)

  // Return a 404 if the template doesn't exist
  info, err := os.Stat(fp)
  if err != nil {
    if os.IsNotExist(err) {
      http.NotFound(w, r)
      return
    }
  }

  // Return a 404 if the request is for a directory
  if info.IsDir() {
    http.NotFound(w, r)
    return
  }

  tmpl, err := template.ParseFiles(lp, fp)
  if err != nil {
    // Log the detailed error
    log.Println(err.Error())
    // Return a generic "Internal Server Error" message
    http.Error(w, http.StatusText(500), 500)
    return
  }

  if err := tmpl.ExecuteTemplate(w, "layout", nil); err != nil {
    log.Println(err.Error())
    http.Error(w, http.StatusText(500), 500)
  }
}

A Recap of Request Handling in Go
=================================

Processing HTTP requests with Go is primarily about two things: ServeMuxes and Handlers.

A ServeMux is essentially a HTTP request router (or multiplexor). It compares
incoming requests against a list of predefined URL paths, and calls the
associated handler for the path whenever a match is found.

Handlers are responsible for writing response headers and bodies. Almost any
object can be a handler, so long as it satisfies the http.Handler interface. In
lay terms, that simply means it must have a ServeHTTP method with the following
signature:

    ServeHTTP(http.ResponseWriter, *http.Request)

Go's HTTP package ships with a few functions to generate common handlers, such
as FileServer, NotFoundHandler and RedirectHandler. Let's begin with a simple
but contrived example:

File: main.go

package main

import (
  "log"
  "net/http"
)

func main() {
  mux := http.NewServeMux()

  rh := http.RedirectHandler("http://example.org", 307)
  mux.Handle("/foo", rh)

  log.Println("Listening...")
  http.ListenAndServe(":3000", mux)
}

The eagle-eyed of you might have noticed something interesting: The signature for the
ListenAndServe function is ListenAndServe(addr string, handler Handler), but we passed a
ServeMux as the second parameter.

We were able to do this because ServeMux also has a ServeHTTP method, meaning that it too
satisfies the Handler interface.

For me it simplifies things to think of a ServeMux as just being a special kind of
handler, which instead of providing a response itself passes the request on to a second
handler. This isn't as much of a leap as it first sounds – chaining handlers together is
fairly commonplace in Go.

Custom Handlers
===============

File: main.go

package main

import (
  "log"
  "net/http"
  "time"
)

type timeHandler struct {
  format string
}

func (th *timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  tm := time.Now().Format(th.format)
  w.Write([]byte("The time is: " + tm))
}

func main() {
  mux := http.NewServeMux()

  th := &timeHandler{format: time.RFC1123}
  mux.Handle("/time", th)

  log.Println("Listening...")
  http.ListenAndServe(":3000", mux)
}

Notice too that we could easily reuse the timeHandler in multiple routes:

func main() {
  mux := http.NewServeMux()

  th1123 := &timeHandler{format: time.RFC1123}
  mux.Handle("/time/rfc1123", th1123)

  th3339 := &timeHandler{format: time.RFC3339}
  mux.Handle("/time/rfc3339", th3339)

  log.Println("Listening...")
  http.ListenAndServe(":3000", mux)
}

Functions as Handlers
=====================

For simple cases (like the example above) defining new custom types and ServeHTTP methods
feels a bit verbose. Let's look at an alternative approach, where we leverage Go's
http.HandlerFunc type to coerce a normal function into satisfying the Handler interface.

Any function which has the signature func(http.ResponseWriter, *http.Request) can be
converted into a HandlerFunc type. This is useful because HandleFunc objects come with an
inbuilt ServeHTTP method which – rather cleverly and conveniently – executes the content
of the original function.

If that sounds confusing, try taking a look at the relevant source code. You'll see that
it's a very succinct way of making a function satisfy the Handler interface.

Let's reproduce the timeHandler application using this technique:

File: main.go

package main

import (
  "log"
  "net/http"
  "time"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
  tm := time.Now().Format(time.RFC1123)
  w.Write([]byte("The time is: " + tm))
}

func main() {
  mux := http.NewServeMux()

  // Convert the timeHandler function to a HandleFunc type
  th := http.HandlerFunc(timeHandler)
  // And add it to the ServeMux
  mux.Handle("/time", th)

  log.Println("Listening...")
  http.ListenAndServe(":3000", mux)
}

In fact, converting a function to a HandlerFunc type and then adding it to a ServeMux like
this is so common that Go provides a shortcut: the ServeMux.HandleFunc method.

This is what the main() function would have looked like if we'd used this shortcut instead:

func main() {
  mux := http.NewServeMux()

  mux.HandleFunc("/time", timeHandler)

  log.Println("Listening...")
  http.ListenAndServe(":3000", mux)
}

Most of the time using a function as a handler like this works well. But there is a bit of
a limitation when things start getting more complex.

You've probably noticed that, unlike the method before, we've had to hardcode the time
format in the timeHandler function. What happens when we want to pass information or
variables from main() to a handler?

A neat approach is to put our handler logic into a closure, and close over the variables
we want to use:

File: main.go

package main

import (
  "log"
  "net/http"
  "time"
)

func timeHandler(format string) http.Handler {
  fn := func(w http.ResponseWriter, r *http.Request) {
    tm := time.Now().Format(format)
    w.Write([]byte("The time is: " + tm))
  }
  return http.HandlerFunc(fn)
}

func main() {
  mux := http.NewServeMux()

  th := timeHandler(time.RFC1123)
  mux.Handle("/time", th)

  log.Println("Listening...")
  http.ListenAndServe(":3000", mux)
}

The timeHandler function now has a subtly different role. Instead of coercing the function
into a handler (like we did previously), we are now using it to return a handler. There's
two key elements to making this work.

First it creates fn, an anonymous function which accesses ‐ or closes over – the format
variable forming a closure. Regardless of what we do with the closure it will always be
able to access the variables that are local to the scope it was created in – which in this
case means it'll always have access to the format variable.

Secondly our closure has the signature func(http.ResponseWriter, *http.Request). As you
may remember from earlier, this means that we can convert it into a HandlerFunc type (so
that it satisfies the Handler interface). Our timeHandler function then returns this
converted closure.

In this example we've just been passing a simple string to a handler. But in a real-world
application you could use this method to pass database connection, template map, or any
other application-level context. It's a good alternative to using global variables, and
has the added benefit of making neat self-contained handlers for testing.

You might also see this same pattern written as:

func timeHandler(format string) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    tm := time.Now().Format(format)
    w.Write([]byte("The time is: " + tm))
  })
}

Or using an implicit conversion to the HandlerFunc type on return:

func timeHandler(format string) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    tm := time.Now().Format(format)
    w.Write([]byte("The time is: " + tm))
  }
}

The DefaultServeMux
===================

You've probably seen DefaultServeMux mentioned in lots of places, from the simplest Hello
World examples to the Go source code.

It took me a long time to realise it isn't anything special. The DefaultServeMux is just a
plain ol' ServeMux like we've already been using, which gets instantiated by default when
the HTTP package is used. Here's the relevant line from the Go source:

var DefaultServeMux = NewServeMux()

The HTTP package provides a couple of shortcuts for working with the DefaultServeMux:
http.Handle and http.HandleFunc. These do exactly the same as their namesake functions
we've already looked at, with the difference that they add handlers to the DefaultServeMux
instead of one that you've created.

Additionally, ListenAndServe will fall back to using the DefaultServeMux if no other
handler is provided (that is, the second parameter is set to nil).

So as a final step, let's update our timeHandler application to use the DefaultServeMux
instead:

File: main.go

package main

import (
  "log"
  "net/http"
  "time"
)

func timeHandler(format string) http.Handler {
  fn := func(w http.ResponseWriter, r *http.Request) {
    tm := time.Now().Format(format)
    w.Write([]byte("The time is: " + tm))
  }
  return http.HandlerFunc(fn)
}

func main() {
  // Note that we skip creating the ServeMux...

  var format string = time.RFC1123
  th := timeHandler(format)

  // We use http.Handle instead of mux.Handle...
  http.Handle("/time", th)

  log.Println("Listening...")
  // And pass nil as the handler to ListenAndServe.
  http.ListenAndServe(":3000", nil)
}

--------------------------------------------------------------------------------
File: main.go

package main

import (
  "html/template"
  "net/http"
  "path"
)

type Profile struct {
  Name    string
  Hobbies []string
}

func main() {
  http.HandleFunc("/", foo)
  http.ListenAndServe(":3000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
  profile := Profile{"Alex", []string{"snowboarding", "programming"}}

  lp := path.Join("templates", "layout.html")
  fp := path.Join("templates", "index.html")

  // Note that the layout file must be the first parameter in ParseFiles
  tmpl, err := template.ParseFiles(lp, fp)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  if err := tmpl.Execute(w, profile); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}
