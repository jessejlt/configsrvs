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
	fmt.Printf("Options = %v\n", cli)
	dataStore, err := GetDataStore(cli.DataStore)
	if err != nil {
		log.Fatal(err)
	}
	dataStore.validate()

	http.HandleFunc("/stats", statsHandler)
	err = http.ListenAndServe(fmt.Sprintf(":%d", cli.Port), nil)
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
	fmt.Printf("Listening on port %d", cli.Port)
}
