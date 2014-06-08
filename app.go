package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/stats", statsHandler)
	http.ListenAndServe(":8080", nil)
	fmt.Printf("Listening on port 8080")
}
