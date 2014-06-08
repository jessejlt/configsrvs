package main

import (
	"os"
	"strings"
)

type Creds struct {
	Region, Key, Secret string
}

func AWSCredentials() (Creds, error) {

	creds := Creds{}
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")

		switch pair[0] {
		case "AWS_DEFAULT_REGION":
			creds.Region = strings.TrimSpace(pair[1])
		case "AWS_ACCESS_KEY_ID":
			creds.Key = strings.TrimSpace(pair[1])
		case "AWS_SECRET_ACCESS_KEY":
			creds.Secret = strings.TrimSpace(pair[1])
		}
	}

	// TODO
	// Check length of creds to ensure they were set,
	// return error if not set

	return creds, nil
}
