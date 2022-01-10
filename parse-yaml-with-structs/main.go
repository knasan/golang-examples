package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// Config all
type config struct {
	Hierarchy []hierarchy
}

type hierarchy struct {
	Name string
	Path string
}

func main() {
	yamlFile, err := ioutil.ReadFile("hierarchy.yaml")

	if err != nil {
		panic(err)
	}

	var c config

	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		panic(err)
	}

	// fmt.Printf("%+v\n", c)

	for _, k := range c.Hierarchy {
		fmt.Printf("Name: %s, Path: %s\n", k.Name, k.Path)
	}
}
