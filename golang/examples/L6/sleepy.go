// http://dougblack.io/words/a-restful-micro-framework-in-go.html
package sleepy

import (
    "encoding/json"
    "fmt"
    "net/http"
    "net/url"
)

const (
    GET    = "GET"
    POST   = "POST"
    PUT    = "PUT"
    DELETE = "DELETE"
)

type Resource interface {
    Get(values url.Values) (int, interface{})
    Post(values url.Values) (int, interface{})
    Put(values url.Values) (int, interface{})
    Delete(values url.Values) (int, interface{})
}

type (
    GetNotSupported    struct{}
    PostNotSupported   struct{}
    PutNotSupported    struct{}
    DeleteNotSupported struct{}
)

func (GetNotSupported) Get(values url.Values) (int, interface{}) {
    return 405, ""
}

func (PostNotSupported) Post(values url.Values) (int, interface{}) {
    return 405, ""
}

func (PutNotSupported) Put(values url.Values) (int, interface{}) {
    return 405, ""
}

func (DeleteNotSupported) Delete(values url.Values) (int, interface{}) {
    return 405, ""
}

type API struct{}

func (api *API) Abort(rw http.ResponseWriter, statusCode int) {
    rw.WriteHeader(statusCode)
}

func (api *API) requestHandler(resource Resource) http.HandlerFunc {
    return func(rw http.ResponseWriter, request *http.Request) {

        var data interface{}
        var code int

        request.ParseForm()
        method := request.Method
        values := request.Form

        switch method {
        case GET:
            code, data = resource.Get(values)
        case POST:
            code, data = resource.Post(values)
        case PUT:
            code, data = resource.Put(values)
        case DELETE:
            code, data = resource.Delete(values)
        default:
            api.Abort(rw, 405)
            return
        }

        content, err := json.Marshal(data)
        if err != nil {
            api.Abort(rw, 500)
        }
        rw.WriteHeader(code)
        rw.Write(content)
    }
}

func (api *API) AddResource(resource Resource, path string) {
    http.HandleFunc(path, api.requestHandler(resource))
}

func (api *API) Start(port int) {
    portString := fmt.Sprintf(":%d", port)
    http.ListenAndServe(portString, nil)
}

package main

import (
    "net/url"
    "sleepy"
)

type HelloResource struct {
    sleepy.PostNotSupported
    sleepy.PutNotSupported
    sleepy.DeleteNotSupported
}

func (HelloResource) Get(values url.Values) (int, interface{}) {
    data := map[string]string{"hello": "world"}
    return 200, data
}

func main() {

    helloResource := new(HelloResource)

    var api = new(sleepy.API)
    api.AddResource(helloResource, "/hello")
    api.Start(3000)

}