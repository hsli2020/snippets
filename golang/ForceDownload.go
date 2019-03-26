package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// change to your target filename
// that you want to force download to your visitor browser
var file = "img.pdf"

func Home(w http.ResponseWriter, r *http.Request) {

	// We will force download after 2 seconds. If you want to increase the delay
	// simply change 2 to whatever number you wanted.

	html := "<html><meta http-equiv='refresh' content='2; url=http://localhost:8080/download' />"
	html = html + "Downloading file now. If the download does not happen automagically. Please "
	html = html + "<a href=" + file + ">click here</a></html>"
	w.Write([]byte(html))
}

func ForceDownload(w http.ResponseWriter, r *http.Request) {

	downloadBytes, err := ioutil.ReadFile(file)

	if err != nil {
		fmt.Println(err)
	}

	// set the default MIME type to send
	mime := http.DetectContentType(downloadBytes)

	fileSize := len(string(downloadBytes))

	// Generate the server headers
	w.Header().Set("Content-Type", mime)
	w.Header().Set("Content-Disposition", "attachment; filename="+file+"")
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Content-Length", strconv.Itoa(fileSize))
	w.Header().Set("Content-Control", "private, no-transform, no-store, must-revalidate")

	//b := bytes.NewBuffer(downloadBytes)
	//if _, err := b.WriteTo(w); err != nil {
	//              fmt.Fprintf(w, "%s", err)
	//      }

	// force it down the client's.....
	http.ServeContent(w, r, file, time.Now(), bytes.NewReader(downloadBytes))
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/download", ForceDownload)

	// SECURITY : Only expose the file permitted for download.
	http.Handle("/"+file, http.FileServer(http.Dir("./")))

	http.ListenAndServe(":8080", nil)
}
