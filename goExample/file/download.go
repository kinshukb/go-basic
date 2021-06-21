package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

// Downloading a File Over HTTP
func main() {
	// Create output file
	newFile, err := os.Create("google.html")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	// HTTP GET request
	url := "http://www.google.com"
	response, err := http.Get(url)
	defer response.Body.Close()

	// Write bytes from HTTP response to file.
	// response.Body satisfies the reader interface.
	// newFile satisfies the writer interface.
	// That allows us to use io.Copy which accepts
	// any type that implements reader and writer interface
	numBytesWritten, err := io.Copy(newFile, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Downloaded %d byte file.\n", numBytesWritten)
}
