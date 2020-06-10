package main

import (
	"fmt"
	"net/http"
	"sort"
	"strings"
)

// formatRequest generates ascii representation of a request
// credit to https://medium.com/doing-things-right/pretty-printing-http-requests-in-golang-a918d5aaa000
func formatRequest(r *http.Request) string {
	// Create return string
	var request []string
	// Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	headerNames := make([]string, 0, len(r.Header))
	for k := range r.Header {
		headerNames = append(headerNames, k)
	}
	// order by header name
	sort.Strings(headerNames)
	// Loop through headers
	for _, name := range headerNames {
		request = append(request, fmt.Sprintf("%v: %v", name, r.Header[name]))
	}

	// If this is a POST, add post data
	if r.Method == "POST" {
		r.ParseForm()
		request = append(request, "\n")
		request = append(request, r.Form.Encode())
	}
	// Return the request as a string
	return strings.Join(request, "\n")
}

func main() {
	var page = `<body>
	<h1>Hello, you've requested: %s</h1>
	<h2>Request body: </h2>
	<pre>%s</pre>
	<body>`

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, page, r.URL.Path, formatRequest(r))
	})

	fmt.Println("Visit http://localhost:9080")
	http.ListenAndServe(":9080", nil)
}
