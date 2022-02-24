package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// maxWorker := 5
	jobs := 10
	ch := make(chan string, jobs)
	wg := &sync.WaitGroup{}

	for i := 1; i <= jobs; i++ {
		ch <- fmt.Sprintf("Job %d", i)
	}
	close(ch)

	// missing jobs 1,3,5,7,9 ?
	for range ch {
		wg.Add(1)
		fmt.Printf("start job: %v\n", <-ch)
		go routine1(ch, wg)
	}
	wg.Wait()

}

// TODO: get information from channel
func routine1(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("routine1 wait 10 seconds")
	time.Sleep(10 * time.Second)
	fmt.Printf("Channel %s\n", <-ch)
	fmt.Printf("%#v\n", ch)
	c := <-ch
	fmt.Println(c)
}
