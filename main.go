package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server running at http://localhost:8932")
	http.ListenAndServe(":8932", nil)
}
