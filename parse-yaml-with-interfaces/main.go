package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

func main() {
	var data interface{}
	b, err := ioutil.ReadFile("input.yaml")
	if err != nil {
		log.Fatal(err)
	}
	if err = yaml.Unmarshal(b, &data); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nData:\n%+v\n", data)

	m, err := yaml.Marshal(&data)
	if err != nil {
		log.Fatal(err)
	}

	if err = ioutil.WriteFile("output.yaml", m, 0644); err != nil {
		log.Fatal(err)
	}
}