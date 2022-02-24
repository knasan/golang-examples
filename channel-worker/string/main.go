package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	maxWorker := 2
	jobs := 10

	wg := &sync.WaitGroup{}
	ch := make(chan string, jobs)

	// start worker
	for i := 1; i <= maxWorker; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range ch {
				routine1(ch, wg)
			}

		}()
	}

	// Insert Jobs
	for i := 0; i <= jobs; i++ {
		ch <- fmt.Sprintf("Job %d", i)
	}
	close(ch)

	wg.Wait()

}

func routine1(ch chan string, wg *sync.WaitGroup) {
	c := <-ch
	fmt.Printf("routine1 for %s wait 10 seconds\n", c)
	time.Sleep(10 * time.Second)
	fmt.Println("Channel", c)
}
