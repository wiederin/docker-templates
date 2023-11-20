package main

// load required packages
import (
	"fmt"
    "strings"

	"github.com/gin-gonic/gin"
    "github.com/gin-contrib/static"
)

func main() {
	// start the server
	serveApplication()
}

func serveApplication() {
	router := gin.Default()

    // serve react frontend
    router.Use(static.Serve("/", static.LocalFile("./frontend/build", true)))

    // no route function
	router.NoRoute(func(c *gin.Context) {
		if !strings.HasPrefix(c.Request.RequestURI, "/api") {
            // if non /api route serve frontend to let react handle routing
			c.File("./frontend/build/index.html")
		}
        // todo: default 404 page not found
	})

	router.Run(":8000")
	fmt.Println("Server running on port 8000")
}
