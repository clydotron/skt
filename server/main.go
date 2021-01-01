package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

type application struct {
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		addr := flag.String("addr", ":8080", "HTTP network address")
		port = *addr
	}

	app := application{}

	fmt.Println("Listening on", port)
	http.Handle("/", http.FileServer(http.Dir("./web")))
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./web"))))
	http.Handle("/api/ping", http.HandlerFunc(app.ping))
	http.ListenAndServe(port, nil)

	// r := gin.Default()
	// // Dont worry about this line just yet, it will make sense in the Dockerise bit!
	// r.Use(static.Serve("/", static.LocalFile("./web", true)))
	// api := r.Group("/api")
	// api.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "poop",
	// 	})
	// })

	// r.Run()
}
