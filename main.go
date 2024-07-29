package main

import (
	"log"
	"net/http"

	"github.com/selvamtech08/equation-compile/handlers"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handlers.TeXGetForm)
	mux.HandleFunc("POST /", handlers.TeXPostForm)
	log.Println("server running on port 8083")
	log.Fatalln(http.ListenAndServe("localhost:8083", mux))
}
