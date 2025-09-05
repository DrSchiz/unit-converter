package main

import (
	"fmt"
	"log"
	"net/http"

	"go-web/handlers"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/length", handlers.Length)
	http.HandleFunc("/weight", handlers.Weight)
	http.HandleFunc("/temperature", handlers.Temperature)

	fmt.Println("Server running on :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
