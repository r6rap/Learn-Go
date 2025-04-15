package webserver

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world")
}

func WebServer() {
	http.HandleFunc("/", index)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Halo")
	})

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}