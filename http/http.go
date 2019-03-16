package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", myController)
	fmt.Println("Serving on http://localhost:8055")
	http.ListenAndServe(":8055", nil)
}

func myController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "What you lookin at?")
}

// go run http.go
