package main

import "flag"

type CLI struct {
	Port   int
	Bucket string
}

func Commands() (CLI, error) {
	cli := CLI{}
	flag.IntVar(&cli.Port, "port", 8080, "Custom port number")
	flag.StringVar(&cli.Bucket, "bucket", "config", "AWS S3 bucket name")
	flag.Parse()

	return cli, nil
}
