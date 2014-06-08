package main

import (
	"errors"
	"os"
	"unicode/utf8"
)

type Creds struct {
	Region, Key, Secret string
}

func AWSCredentials() (Creds, error) {

	creds := Creds{}
	creds.Region = os.Getenv("AWS_DEFAULT_REGION")
	creds.Key = os.Getenv("AWS_ACCESS_KEY_ID")
	creds.Secret = os.Getenv("AWS_SECRET_ACCESS_KEY")

	if utf8.RuneCountInString(creds.Region) == 0 {
		return creds, errors.New("Missing required env.AWS_DEFAULT_REGION")
	}

	if utf8.RuneCountInString(creds.Key) == 0 {
		return creds, errors.New("Missing required env.AWS_ACCESS_KEY_ID")
	}

	if utf8.RuneCountInString(creds.Secret) == 0 {
		return creds, errors.New("Missing required env.AWS_SECRET_ACCESS_KEY")
	}

	return creds, nil
}
