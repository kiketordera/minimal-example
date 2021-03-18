package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// We create the instance for Gin
	r := gin.Default()

	// Path to the static files. /static is rendered in the HTML and /media is the link to the path to the  images, svg, css.. the static files
	r.StaticFS("/static", http.Dir("../media"))

	// Path to the HTML templates. * is a wildcard
	r.LoadHTMLGlob("../media/html/*/*.html")

	r.NoRoute(renderHome)

	// This get executed when the users gets into our website in the home domain ("/")
	r.GET("/", renderHome)

	r.Run(":8080")

}

func renderHome(c *gin.Context) {
	c.HTML(http.StatusOK, "my-html.html", gin.H{})
}
