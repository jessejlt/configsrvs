package main

import (
	"fmt"
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
	fmt.Printf("Using AWS Creds = %v\n", creds)
	s.Credentials = creds

	return true
}

func (s S3) fetch() (string, error) {
	return "", nil
}

func (s S3) update(data string) error {
	return nil
}
