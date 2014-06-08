package main

import "testing"

func TestEmptyArgs(t *testing.T) {

	const port = 8080
	const bucket = "config"
	const dataStore = "s3"
	cli, err := Commands()
	if err != nil {
		t.Error(err)
	}

	if port != cli.Port {
		t.Error("Expected %v, Received %v", port, cli.Port)
	}

	if bucket != cli.Bucket {
		t.Error("Expected %v, Received %v", bucket, cli.Bucket)
	}

	if dataStore != cli.DataStore {
		t.Error("Expected %v, Received %v", dataStore, cli.DataStore)
	}
}
