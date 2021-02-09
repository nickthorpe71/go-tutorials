package main

import (
	"log"
	"net/http"
)

// Tutorial
// https://www.youtube.com/watch?v=VzBGi_n65iU&ab_channel=NicJackson

func main() {
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		log.Println("Hello Internet")
	})

	http.ListenAndServe(":9090", nil)
}
