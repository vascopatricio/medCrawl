package main

import (
	"net/http"
)

func main() {
  //	Frontend - HTML
	http.HandleFunc("/", handler)

  //	Backend - API
	http.HandleFunc("/api/topics/", topicsResource)
	http.HandleFunc("/api/articles/", articlesResource)

	http.ListenAndServe(":3000", nil)
}