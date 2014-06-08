package main

import (
  "io/ioutil"
  "log"
)

type S3 struct {
  Credentials Creds
  ETag        string
}

func (s S3) validate() bool {

  creds, err := AWSCredentials()
  if err != nil {
    log.Fatal(err)
    return false
  }
  s.Credentials = creds

  return true
}

func (s S3) fetch() (string, error) {

  data, err := ioutil.ReadFile("config.json")
  if err != nil {
    return "", err
  }

  return string(data), nil
}

func (s S3) update(data string) error {
  return nil
}
