package main

import (
	"fmt"
	"net/http"
)

const port = 3000

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})

	http.HandleFunc("/feed", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(feeds.generateCombinedFeed())
	})

	http.ListenAndServe(":"+port, nil)
}
