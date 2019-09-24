package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/davidbyttow/govips"
)

func main() {
	// Start vips with the default configuration
	vips.Startup(nil)
	defer vips.Shutdown()

	http.HandleFunc("/", resizeHandler)
	log.Fatal(http.ListenAndServe("0.0.0.0:4444", nil))
}

func resizeHandler(w http.ResponseWriter, r *http.Request) {
	// Get the query parameters from the request URL
	query := r.URL.Query()
	queryUrl := query.Get("url")
	queryWidth := query.Get("width")
	queryHeight := query.Get("height")

	// Validate that all three required fields are present
	if queryUrl == "" || queryWidth == "" || queryHeight == "" {
		w.Write([]byte(fmt.Sprintf("url, width and height are required")))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Convert width and height to integers
	width, errW := strconv.Atoi(queryWidth)
	height, errH := strconv.Atoi(queryHeight)
	if errW != nil || errH != nil {
		w.Write([]byte(fmt.Sprintf("width and height must be integers")))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Start fetching the image from the given url
	resp, err := http.Get(queryUrl)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("failed to get %s: %v", queryUrl, err)))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Ensure that a valid response was given
	if resp.StatusCode/100 != 2 {
		w.Write([]byte(fmt.Sprintf("failed to get %s: status %d", queryUrl, resp.StatusCode)))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer resp.Body.Close()

	// govips returns the output of the image as a []byte object. We don't need
	// it since we are directly piping it to the ResponseWriter
	_, err = vips.NewTransform().
		Load(resp.Body).
		ResizeStrategy(vips.ResizeStrategyStretch).
		Resize(width, height).
		Output(w).
		Apply()

	if err != nil {
		w.Write([]byte(fmt.Sprintf("failed to resize %s: %v", queryUrl, err)))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}