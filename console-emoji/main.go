package main

import (
	"fmt"
	animals "github.com/hackebrot/turtle"
	"os"
)

const animal = "penguin"

func main() {

	emoji, ok := animals.Emojis[animal]
	if !ok {
		fmt.Fprintf(os.Stderr, "Char %v not found\n", animal)
		os.Exit(1)
	}

	fmt.Println(emoji.Char)
}
