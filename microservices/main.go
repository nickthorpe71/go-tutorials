package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Tutorial
// 1. https://www.youtube.com/watch?v=VzBGi_n65iU&ab_channel=NicJackson
// 2. https://www.youtube.com/watch?v=hodOppKJm5Y&ab_channel=NicJackson <-

// GO HTTP docs
// https://golang.org/pkg/net/http/

func main() {
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		log.Println("Hello Internet")
		data, err := ioutil.ReadAll(request.Body)
		if err != nil {
			http.Error(response, "Oooops", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(response, "Hello %s\n", data)
	})

	http.HandleFunc("/north", func(http.ResponseWriter, *http.Request) {
		log.Println("after heading north you see a cliff")
	})

	http.ListenAndServe(":9090", nil)
}
