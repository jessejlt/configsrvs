package main

import "encoding/json"

type jsonResponse map[string]interface{}

func (r jsonResponse) String() (s string) {
	b, err := json.Marshal(r)
	if err != nil {
		s = ""
		return
	}
	s = string(b)
	return
}
