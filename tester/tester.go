package main

import (
	"example.com/practice"
	"example.com/library"
	"fmt"
	"net/http"
	"log"
)

func main() {
	practice.Test()
	http.HandleFunc("/", library.HelloWorld)
	fmt.Println("Server started and listening on localhost:9093")
	log.Fatal(http.ListenAndServe(":9003", nil))
}