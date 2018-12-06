package webservice

import (
	"net/http"
	"github.com/codegangsta/martini"
)

// WebService is the interface that should be implemented by types that want to
// provide web services.
type WebService interface {
	// GetPath returns the path to be associated with the service.
	GetPath() string

	// WebDelete wraps a DELETE method request. The given params might be
	// empty, in case it was applied to the collection itself (i.e. all
	// entries instead of a single one) or will have a "id" key that will
	// point to the id of the entry being deleted.
	WebDelete(params martini.Params) (int, string)

	// WebGet is Just as above, but for the GET method. If params is empty,
	// it returns all the entries in the collection. Otherwise it returns
	// the entry with the id as per the "id" key in params.
	WebGet(params martini.Params) (int, string)

	// WebPost wraps the POST method. Again an empty params means that the
	// request should be applied to the collection. A non-empty param will
	// have an "id" key that refers to the entry that should be processed
	// (note this specific case is usually not supported unless each entry
	// is also a collection).
	WebPost(params martini.Params, req *http.Request) (int, string)
}

// RegisterWebService adds martini routes to the relevant webservice methods
// based on the path returned by GetPath. Each method is registered once for
// the collection and once for each id in the collection.
func RegisterWebService(webService WebService,
	classicMartini *martini.ClassicMartini) {
	path := webService.GetPath()

	classicMartini.Get(path, webService.WebGet)
	classicMartini.Get(path+"/:id", webService.WebGet)

	classicMartini.Post(path, webService.WebPost)
	classicMartini.Post(path+"/:id", webService.WebPost)

	classicMartini.Delete(path, webService.WebDelete)
	classicMartini.Delete(path+"/:id", webService.WebDelete)
}
package main

import (
	"github.com/brunoga/go-webservice-sample/guestbook"
	"github.com/brunoga/go-webservice-sample/webservice"
	"github.com/codegangsta/martini"
)
func main() {
	martiniClassic := martini.Classic()
	guestBook := guestbook.NewGuestBook()
	webservice.RegisterWebService(guestBook, martiniClassic)
	martiniClassic.Run()
}
