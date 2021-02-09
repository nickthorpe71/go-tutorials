package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Tutorial
// https://www.youtube.com/watch?v=VzBGi_n65iU&ab_channel=NicJackson

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello Internet")
		data, _ := ioutil.ReadAll(r.Body)
		fmt.Fprintf(rw, "Hello %s\n", data)
	})

	http.HandleFunc("/north", func(http.ResponseWriter, *http.Request) {
		log.Println("after heading north you see a cliff")
	})

	http.ListenAndServe(":9090", nil)
}
