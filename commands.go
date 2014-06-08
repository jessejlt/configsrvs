package main

import "flag"

type CLI struct {
	Port int
}

func Commands() (CLI, error) {
	cli := CLI{}
	flag.IntVar(&cli.Port, "port", 8080, "Add custom port number")
	flag.Parse()

	return cli, nil
}
