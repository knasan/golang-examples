package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan string, 5)
	wg.Add(4)
	go routine1(ch, wg)
	go routine2(ch, wg)
	go routine3(ch, wg)
	go routine3(ch, wg)
	wg.Wait()
}

// routine1 is the sender
func routine1(ch chan string, wg *sync.WaitGroup) {
	fmt.Println(("routine1 wait 10 seconds"))
	time.Sleep(10 * time.Second)
	ch <- "hello from routine1!"
	wg.Done()
}

// routine2 is a other sender
func routine2(ch chan string, wg *sync.WaitGroup) {
	fmt.Println("routine2 wait 4 seconds")
	time.Sleep(4 * time.Second)
	ch <- "important message from routine2!"
	wg.Done()
}

// routine3 is the receiver
func routine3(ch chan string, wg *sync.WaitGroup) {
	// wait for message
	message := <-ch
	fmt.Println("Message: ", message)
	wg.Done()
}