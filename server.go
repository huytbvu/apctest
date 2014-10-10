package main

import (
	"log"
	"net/http"
)

func init() {
	log.Println("Starting Apptest")
}

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello World"))
	})
	http.ListenAndServe(":4040", nil)
}
