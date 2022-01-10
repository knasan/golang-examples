package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Destination string
	MaxWorker   int
	Default     Default
	Hosts       map[string]Host
}

type Host struct {
	Address string
	Port    int
	Plugins map[string]Plugin `yaml:"plugins"`
}

type Plugin struct {
	Include string
	Exclude []string
	// before,after and recovery commands
	AfterCommands   map[string]Command `yaml:"afterCommands"`
	BeforeCommands  map[string]Command `yaml:"beforeCommands"`
	RestoreCommands map[string]Command `yaml:"restoreCommands"`
}

type Command struct {
	Name            string
	Content         []string
	ContinueOnError bool `yaml:"continueOnError"`
}

type Default struct {
	MinSpaceLeft string   `yaml:"minSpaceLeft"`
	MaxSnapshots int      `yaml:"maxSnapshots"`
	RsyncOptions []string `yaml:"rsyncOptions"`
}

func main() {
	var data Config
	// read file as []byte
	b, err := ioutil.ReadFile("config.yaml")
	// check error for ReadFile
	if err != nil {
		log.Fatalf("Error load yamlfile: %+v\n", err)
	}

	// yaml unmarshal and check error
	if err = yaml.Unmarshal(b, &data); err != nil {
		log.Fatal(err)
	}

	// test output
	fmt.Printf("Config: %+v\n", data)

	// check error
	m, err := yaml.Marshal(&data)
	if err != nil {
		log.Fatal(err)
	}

	// write file
	if err = ioutil.WriteFile("output.yaml", m, 0644); err != nil {
		log.Fatal(err)
	}
}
