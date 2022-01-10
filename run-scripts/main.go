package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
)

func main() {
	files, err := ioutil.ReadDir("scripts")

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fileName := file.Name()
		if err = run("scripts", fileName); err != nil {
			fmt.Println(err)
		}
	}
}

func run(path, file string) (err error) {
	fmt.Println("path:", path, "file:", file)

	cmd := exec.Command(filepath.Join(path, file))

	//cmd.Stderr = os.Stderr
	//cmd.Stdout = os.Stdout

	if err = cmd.Start(); err != nil {
		return err
	}

	if err = cmd.Wait(); err != nil {
		return err
	}

	return nil
}
