package main

import (
	"github.com/hatzelencio/merge-branch/remote"
	"log"
)

func main() {
	var err error

	err = remote.ValidateInputs()

	if err != nil {
		log.Fatal(err)
	}

	err = remote.Merge()

	if err != nil {
		log.Fatal(err)
	}
}
