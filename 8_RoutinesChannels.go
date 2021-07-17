package main

import (
	"fmt"
	"runtime"
)

/*
- lightweight threads, no id of go routines
- not constant time of execution
- not always paralelism - scheduler, cores, processors
- m:n planning -> m go routine on n threads of operation system
- GOMAXPROCS == m parameter in m:n planning
- executor tracing
- main go routine terminates child goroutines
- channnels = sync and data flow beetween goroutines
- channels (FIFO), channel has pointer under the hood, nil,
work with any data type
- channel types: bideractional, singledirectional,
bufferized [cap/len functions] or not, open/closed
- channel operations:
create: ch = make (chan int) || ch = make(chan int), 3)
send to channel: ch <- x
receive from channel: x = <-ch || x,ok (channel opened or closed) = <-ch
close: close(ch), will return nil/0 on read from close channel,
write to closed channel triggers panic, who creates -> closes
- select
- timer
- ticker == cron prog
- interrupt <-interrupt => gracefull shutdown

*/
func main() {
	go fmt.Println("Channels")

	/* quantity of goroutines */
	fmt.Println("Goroutines", runtime.NumGoroutine())

	/* not bufferized channels */
	var ch = make(chan int)

	go func() {
		fmt.Printf("Hello\n")
		ch <- 1
	}()

	/* blocked until received */
	<-ch
}
