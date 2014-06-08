package main

import (
  "encoding/json"
  "html/template"
  "net/http"
  "path"
  "strings"
)

type Home struct {
  Message string
	Config string
}

func serveJSON(w http.ResponseWriter, r *http.Request, h Home) {

  js, err := json.Marshal(h)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
}

func serveHTML(w http.ResponseWriter, r *http.Request, h Home) {

  fp := path.Join("templates", "index.html")
  tmpl, err := template.ParseFiles(fp)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  if err := tmpl.Execute(w, h); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

  RequestCount.Add(1)
	config, err := Store.fetch()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

  home := Home{Message: "TODO", Config: config}
  acceptHeader := r.Header.Get("Accept")
  if strings.Contains(acceptHeader, "application/json") {
    serveJSON(w, r, home)
  } else {
    serveHTML(w, r, home)
  }
}
