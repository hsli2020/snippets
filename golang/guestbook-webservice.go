// WebService related methods.

package guestbook

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/codegangsta/martini"
)

// GetPath implements webservice.GetPath.
func (g *GuestBook) GetPath() string {
	// Associate this service with http://host:port/guestbook.
	return "/guestbook"
}

// WebDelete implements webservice.WebDelete.
func (g *GuestBook) WebDelete(params martini.Params) (int, string) {
	if len(params) == 0 {
		// No params. Remove all entries from collection.
		g.RemoveAllEntries()

		return http.StatusOK, "collection deleted"
	}

	// Convert id to an integer.
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		// Id was not a number.
		return http.StatusBadRequest, "invalid entry id"
	}

	// Remove entry identified by id.
	err = g.RemoveEntry(id)
	if err != nil {
		// Could not find entry.
		return http.StatusNotFound, "entry not found"
	}

	return http.StatusOK, "entry deleted"
}

// WebGet implements webservice.WebGet.
func (g *GuestBook) WebGet(params martini.Params) (int, string) {
	if len(params) == 0 {
		// No params. Return entire collection encoded as JSON.
		encodedEntries, err := json.Marshal(g.GetAllEntries())
		if err != nil {
			// Failed encoding collection.
			return http.StatusInternalServerError, "internal error"
		}

		// Return encoded entries.
		return http.StatusOK, string(encodedEntries)
	}

	// Convert id to integer.
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		// Id was not a number.
		return http.StatusBadRequest, "invalid entry id"
	}

	// Get entry identified by id.
	entry, err := g.GetEntry(id)
	if err != nil {
		// Entry not found.
		return http.StatusNotFound, "entry not found"
	}

	// Encode entry in JSON.
	encodedEntry, err := json.Marshal(entry)
	if err != nil {
		// Failed encoding entry.
		return http.StatusInternalServerError, "internal error"
	}

	// Return encoded entry.
	return http.StatusOK, string(encodedEntry)
}

// WebPost implements webservice.WebPost.
func (g *GuestBook) WebPost(params martini.Params, req *http.Request) (int, string) {

	// Make sure Body is closed when we are done.
	defer req.Body.Close()

	// Read request body.
	requestBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return http.StatusInternalServerError, "internal error"
	}

	if len(params) != 0 {
		// No keys in params. This is not supported.
		return http.StatusMethodNotAllowed, "method not allowed"
	}

	// Unmarshal entry sent by the user.
	var guestBookEntry GuestBookEntry
	err = json.Unmarshal(requestBody, &guestBookEntry)
	if err != nil {
		// Could not unmarshal entry.
		return http.StatusBadRequest, "invalid JSON data"
	}

	// Add entry provided by the user.
	g.AddEntry(guestBookEntry.Email, guestBookEntry.Title, guestBookEntry.Content)

	// Everything is fine.
	return http.StatusOK, "new entry created"
}
