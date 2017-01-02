package main

import (
	"bytes"
	"image"
	"image/png"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/test.png", responseTestPng)

	router.Run(":" + port)
}

func responseTestPng(c *gin.Context) {
	imagefile, err := os.Open("./images/happi_coat.png")
	if err != nil {
		// TODO: Error Handling
		panic(err.Error())
	}
	defer imagefile.Close()
	img, _, err := image.Decode(imagefile)
	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, img); err != nil {
		// TODO: Error Hangling
		panic(err.Error())
	}

	c.Data(http.StatusOK, "image/png", buffer.Bytes())
}
