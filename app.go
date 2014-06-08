package main

import (
  "expvar"
  "fmt"
  "github.com/codegangsta/negroni"
  "log"
  "net/http"
)

var (
  RequestCount = expvar.NewInt("request_count")
  Store        DataStore
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
  Store = dataStore

  mux := http.NewServeMux()
  mux.HandleFunc("/", HomeHandler)

  n := negroni.Classic()
  n.UseHandler(mux)
  n.Run(fmt.Sprintf(":%d", cli.Port))
}
