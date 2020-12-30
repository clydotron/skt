package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

type application struct {
}

func main() {

	// app := application{}

	// addr := flag.String("addr", ":8080", "HTTP network address")

	// fmt.Println("Listening on", *addr)
	// http.ListenAndServe(*addr, app.routes())

	r := gin.Default()
	// Dont worry about this line just yet, it will make sense in the Dockerise bit!
	r.Use(static.Serve("/", static.LocalFile("./web", true)))
	api := r.Group("/api")
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "poopy",
		})
	})

	r.Run()
}
