package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	maxWorker := 5
	jobs := 10

	wg := &sync.WaitGroup{}
	ch := make(chan string)

	for i := 1; i <= maxWorker; i++ {
		go func() {
			routine1(ch, wg)
		}()
	}

	for i := 1; i <= jobs; i++ {
		ch <- fmt.Sprintf("Job %d", i)
	}

	wg.Wait()

}

func routine1(ch chan string, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	fmt.Println(("routine1 wait 10 seconds"))
	time.Sleep(10 * time.Second)
	msg := <-ch
	fmt.Println("Channel", msg)
}
