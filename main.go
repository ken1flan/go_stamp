package main

import (
	"bytes"
	"image"
	"image/png"
	"log"
	"net/http"
	"os"
	"strconv"
)

func response(rw http.ResponseWriter, request *http.Request) {
	rw.Write([]byte("Hello, world."))
}

func responsePng(rw http.ResponseWriter, request *http.Request) {
	imagefile, err := os.Open("/Users/ken1flan/src/go_stamp/happi_coat.png")
	if err != nil {
		// TODO: Error Hangling
		panic(err.Error())
	}
	defer imagefile.Close()
	img, _, err := image.Decode(imagefile)
	if err != nil {
		// TODO: Error Hangling
		panic(err.Error())
	}

	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, img); err != nil {
		log.Println("unable to encode image.")
	}

	rw.Header().Set("Content-Type", "image/jpeg")
	rw.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := rw.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}

func main() {
	http.HandleFunc("/", response)
	http.HandleFunc("/test.png", responsePng)
	http.ListenAndServe(":3000", nil)
}
