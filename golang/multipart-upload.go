package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"time"
)

func upload() error {
	// Metadata content.
	metadata := `{"name": "photo-sample.jpeg"}`

	// New empty buffer
	body := &bytes.Buffer{}
	// Creates a new multipart Writer with a random boundary, writing to the empty
	// buffer
	writer := multipart.NewWriter(body)

	// Metadata part
	metadataHeader := textproto.MIMEHeader{}
	// Set the Content-Type header
	metadataHeader.Set("Content-Type", "application/json; charset=UTF-8")
	// Create new multipart part
	part, err := writer.CreatePart(metadataHeader)
	if err != nil {
		return err
	}
	// Write the part body
	part.Write([]byte(metadata))

	// Media part
	// Read the file to memory
	mediaData, err := ioutil.ReadFile("file.jpeg")
	if err != nil {
		return err
	}
	mediaHeader := textproto.MIMEHeader{}
	mediaHeader.Set("Content-Type", "image/jpeg")

	mediaPart, err := writer.CreatePart(mediaHeader)
	if err != nil {
		return err
	}
	io.Copy(mediaPart, bytes.NewReader(mediaData))

	// Finish the multipart message
	writer.Close()

	accessToken := "<google drive access token>"

	endpoint := "https://www.googleapis.com/upload/drive/v3/files?uploadType=multipart"
	req, err := http.NewRequest("POST", endpoint, bytes.NewReader(body.Bytes()))
	if err != nil {
		return err
	}

	// Request Content-Type with boundary parameter.
	contentType := fmt.Sprintf("multipart/related; boundary=%s", writer.Boundary())
	req.Header.Set("Content-Type", contentType)
	// Content-Length must be the total number of bytes in the request body.
	req.Header.Set("Content-Length", fmt.Sprintf("%d", body.Len()))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	client := &http.Client{Timeout: 120 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	// Do what you want with the response

	return nil
}
