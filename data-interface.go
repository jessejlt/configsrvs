package main

import "errors"

type DataStore interface {
	validate() bool
	fetch() (string, error)
	update(string) error
}

func GetDataStore(storeName string) (DataStore, error) {

	var dataStore DataStore
	switch storeName {
	case "s3":
		dataStore = S3{}
	default:
		dataStore = S3{}
	}

	err := errors.New("DataStore failed validation")
	if dataStore.validate() {
		err = nil
	}

	return dataStore, err
}
