package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
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

	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)

	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(fmt.Sprintf(":%d", cli.Port))
}
