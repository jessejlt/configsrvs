package main

import "flag"

type CLI struct {
	Port              int
	Bucket, DataStore string
}

func Commands() (CLI, error) {
	cli := CLI{}
	flag.IntVar(&cli.Port, "port", 8080, "Custom port number")
	flag.StringVar(&cli.Bucket, "bucket", "config", "AWS S3 bucket name")
	flag.StringVar(&cli.DataStore, "data-store", "s3", "Data backing")
	flag.Parse()

	return cli, nil
}
