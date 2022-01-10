package main

import (
	"log"
	"os"
)

func main() {
	of := "./file"
	nf := "./new-filename"

	if err := os.Rename(of, nf); err != nil {
		log.Fatal(err)
	}
}
