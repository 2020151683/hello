package main

import (
	"log"
	"net/http"

	"github.com/jalpuche/hello/handlers"
)

func main() {
	//create a multiplexer
	mux := http.NewServeMux()
	//create a file server
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/resource/", http.StripPrefix("/resource", fileServer))

	mux.HandleFunc("/greeting", handlers.Greeting)
	mux.HandleFunc("/home", handlers.Home)
	mux.HandleFunc("/about", handlers.About)
	mux.HandleFunc("/message/create", handlers.MessageCreate)

	//create our server
	log.Print("starting server on port :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
