package main

import "fmt"
import "sort"

func main() {
	slice := []string{"plugin01", "plugin02", "plugin03", "plugin10", "plugin23", "plugin11", "plugin22", "001_plugin", "004_plugin", "009_plugin", "007_plugin"}

	fmt.Println(slice)
	sort.Strings(slice)
	fmt.Println(slice)

	for _, ent := range slice {
		fmt.Println(ent)
	}
}
