package main

import "fmt"

type cat string
type dog string
type mouse string

func main() {
	c := cat("Cat")
	d := dog("Dog")
	m := mouse("Mouse")

	var r string
	r = determine(c)
	fmt.Printf("C is %s\n", r)
	r = determine(d)
	fmt.Printf("D is %s\n", r)
	r = determine(m)
	fmt.Printf("M is %s\n", r)
}

func determine(i interface{}) string {
	switch i.(type) {
	case cat:
		return "Cat"
	case dog:
		return "Dog"
	case mouse:
		return "Mouse"
	default:
		return "-"
	}
}
