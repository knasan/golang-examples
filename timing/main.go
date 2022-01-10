package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	t := time.Now()
	t = t.UTC()

	fmt.Println("Current time:", t.Local())

	fmt.Println("wait 20 seconds ...")
	time.Sleep(20 * time.Second)

	endTime := time.Since(t).Seconds()

	fmt.Println(math.Round(endTime))
}
