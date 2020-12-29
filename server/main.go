package main

import (
	"flag"
	"fmt"
	"net/http"
)

type application struct {
}

func main() {

	app := application{}

	addr := flag.String("addr", ":4000", "HTTP network address")

	fmt.Println("Listening on", *addr)
	http.ListenAndServe(*addr, app.routes())

}
