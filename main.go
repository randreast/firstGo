package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ini index")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hallo")
	})

	http.HandleFunc("/index", index)

	fmt.Println("Starting web server")
	http.ListenAndServe(":8000", nil)
}
