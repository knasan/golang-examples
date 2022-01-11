package main

import (
	"fmt"
	"github.com/postfinance/single"
	"time"
)

func main() {
	one, err := single.New("gotestlock", single.WithLockPath("/tmp/"))
	if err != nil {
		panic(err)
	}

	// lock and defer unlocking
	if err := one.Lock(); err != nil {
		panic(err)
	}

	fmt.Println("Start Program")

	time.Sleep(60 * time.Second)

	if err := one.Unlock(); err != nil {
		panic(err)
	}

}
