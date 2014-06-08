package main

import (
	"os"
	"testing"
)

func TestRequiresCreds(t *testing.T) {

	os.Setenv("AWS_DEFAULT_REGION", "")
	os.Setenv("AWS_ACCESS_KEY_ID", "")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "")

	_, err := AWSCredentials()
	if err == nil {
		t.Error("Failed to detect missing AWS credentials")
	}
}

func TestCreds(t *testing.T) {

	const a, b, c = "a", "b", "c"
	os.Setenv("AWS_DEFAULT_REGION", a)
	os.Setenv("AWS_ACCESS_KEY_ID", b)
	os.Setenv("AWS_SECRET_ACCESS_KEY", c)

	creds, err := AWSCredentials()
	if err != nil {
		t.Errorf("Error parsing env vars %v", err)
	}

	if creds.Region != a {
		t.Errorf("AWS_DEFAULT_REGION expected %v received %v", a, creds.Region)
	}

	if creds.Key != b {
		t.Errorf("AWS_ACCESS_KEY_ID expected %v received %v", b, creds.Key)
	}

	if creds.Secret != c {
		t.Errorf("AWS_SECRET_ACCESS_KEY expected %v received %v", c, creds.Secret)
	}
}
