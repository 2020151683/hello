//test

package handlers

import (
	"net/http"
)

func Greeting(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to my page"))
}

// -> home
func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("this is my home page"))
}

// -> about
func About(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is my about page"))
}
func MessageCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("ALlow", "POST")
		w.WriteHeader(405)
		w.Write([]byte("method not allowed"))
		return
	}
	w.Write([]byte("message created..."))
}
