package main

import (
	"net/http"
	"log"
)

func main() {
	fs := http.FileServer(http.Dir("templates/"))
  http.Handle("/", fs)
	log.Println("Server started at localhost:9090")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
