package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	cli, err := Commands()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Port=%d\n", cli.Port)

	http.HandleFunc("/stats", statsHandler)
	http.ListenAndServe(fmt.Sprintf(":%d", cli.Port), nil)
	fmt.Printf("Listening on port %d", cli.Port)
}
