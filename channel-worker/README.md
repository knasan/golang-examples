# channel-worker

is an experiment to define a number of workers. Push all tasks into the worker and it should then process the tasks. As soon as a process becomes free, it should process the next task until there are no more tasks.

I implemented the solution as follows.
For example, if you want to have two workers, then the goroutine must start as often as the number of workers. This requires two simple for loops.

```go
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
```

the outer goroutine is run through as many times as you want to have workers. In this example two workers.
A wait group add is performed for each worker.

within this loop i start the actual worker as a goroutine function, which is a loop over the channel.

Now you have two workers who can process the messages in the channel, as soon as one worker has finished processing, the next message is processed until there are no more messages. As soon as the inner loop (loop over the channel) is exited, a wg.Done() is executed and the goroutine ends.

It may be that a loop with a null message is executed, you should react to that. In this example I have left it so that the loop is seen to execute once more than there are messages.
