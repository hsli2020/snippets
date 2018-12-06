// formatRequest generates ascii representation of a request
func formatRequest(r *http.Request) string {
 // Create return string
 var request []string

 // Add the request string
 url := fmt.Sprintf(“%v %v %v”, r.Method, r.URL, r.Proto)
 request = append(request, url)

 // Add the host
 request = append(request, fmt.Sprintf(“Host: %v”, r.Host))

 // Loop through headers
 for name, headers := range r.Header {
   name = strings.ToLower(name)
   for _, h := range headers {
     request = append(request, fmt.Sprintf(“%v: %v”, name, h))
   }
 }
 
 // If this is a POST, add post data
 if r.Method == “POST” {
    r.ParseForm()
    request = append(request, “\n”)
    request = append(request, r.Form.Encode())
 } 

  // Return the request as a string
  return strings.Join(request, “\n”)
}

the httputil package has a DumpRequest method pre-baked. The output is nearly exactly the same 
which is great since it removes the dependency on 3rd party code.


// Save a copy of this request for debugging.
requestDump, err := httputil.DumpRequest(req, true)
if err != nil {
  fmt.Println(err)
}
fmt.Println(string(requestDump))