package main

import (
	"fmt"
	"net/http"

	"github.com/jwonsever/news/feeds"
)

const port = 3000

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fmt.Fprintf(w, "Welcome to my website!")
	})

	http.HandleFunc("/feed", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		f, _ := feeds.GenerateCombinedFeed()
		fmt.Fprintf(w, f)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
