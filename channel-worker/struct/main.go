package main

import (
	"fmt"
	"sync"
	"time"
)

type Message struct {
	Text  string
	Text2 string
	Text3 string
}

func main() {
	maxWorker := 2
	jobs := 10

	wg := &sync.WaitGroup{}
	ch := make(chan Message, jobs)

	// start worker
	for i := 1; i <= maxWorker; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for c := range ch {
				routine1(c)
			}

		}()
	}

	// Insert Jobs
	for i := 0; i <= jobs; i++ {
		msg := Message{}
		msg.Text = fmt.Sprintf("Hello from Message1 from Job %d", i)
		msg.Text2 = fmt.Sprintf("Hello from Message2 from Job %d", i)
		msg.Text3 = fmt.Sprintf("Hello from Message3 from Job %d", i)
		ch <- msg
	}
	close(ch)

	wg.Wait()

}

func routine1(c Message) {
	fmt.Printf("routine1 for %s wait 10 seconds\n", c)
	time.Sleep(1 * time.Second)
	fmt.Println("Message1: ", c.Text)
	fmt.Println("Message2:", c.Text2)
	fmt.Println("Message3:", c.Text3)
}
