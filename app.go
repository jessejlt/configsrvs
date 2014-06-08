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

	creds, err := AWSCredentials()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("AWS Creds = %v\n", creds)
	fmt.Printf("Options = %v\n", cli)

	http.HandleFunc("/stats", statsHandler)
	err = http.ListenAndServe(fmt.Sprintf(":%d", cli.Port), nil)
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
	fmt.Printf("Listening on port %d", cli.Port)
}
