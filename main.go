package main

import (
	"log"
	"os"
)

func main() {
	// read config from os.Args
	if len(os.Args) <= 1 {
		log.Fatal("can't read config file")
	}

	// TODO:
}
