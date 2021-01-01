package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type application struct {
}

func registerPing(db *sql.DB) {

	_, err := db.Exec("INSERT INTO ping_timestamp (occurred) VALUES ($1)", time.Now())
	if err != nil {
		log.Println("Couldn't insert the ping")
		log.Println(err)
	}
}

func pingFunc(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {

		defer registerPing(db)
		r := db.QueryRow("SELECT occurred FROM ping_timestamp ORDER BY id DESC LIMIT 1")
		var lastDate pq.NullTime
		r.Scan(&lastDate)

		message := "first time!"
		if lastDate.Valid {
			message = fmt.Sprintf("%v ago", time.Now().Sub(lastDate.Time).String())
		}

		fmt.Println(message)
		c.JSON(200, gin.H{
			"message": message,
		})
	}
}

func main() {

	// port := os.Getenv("PORT")
	// fmt.Println("env PORT =", port)
	// if port == "" {
	// 	addr := flag.String("addr", ":8080", "HTTP network address")
	// 	port = *addr
	// }

	// app := application{}

	// fmt.Println("Listening on", port)
	// http.Handle("/", http.FileServer(http.Dir("./web")))
	// http.Handle("/api/ping", http.HandlerFunc(app.ping))
	// http.ListenAndServe(port, nil)

	r := gin.Default()

	// Dont worry about this line just yet, it will make sense in the Dockerise bit!
	r.Use(static.Serve("/", static.LocalFile("./web", true)))
	api := r.Group("/api")

	dbUrl := os.Getenv("DATABASE_URL")
	log.Printf("DB [%s]", dbUrl)
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	api.GET("/ping", pingFunc(db))

	r.Run()
}
