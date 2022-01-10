package main

import "fmt"
import "os"

func main() {
	_, err := os.Stat("filelink")
	if err != nil {
		fmt.Println("not exists")
	} else {
		fmt.Println("exists")
	}
}
