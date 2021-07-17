package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*
- GOMAXPROC > 1
- waitgroup: add[add(-1)], done, wait
- mutex
- GOMAXPROC -> env == go env
- go run/build -race == check race conditions
- 8128 limit of goroutines
- defer for multiple unlocks
- rwMutex == all can read while on write blocks all readers
- sync.Map
- sync.Pool == storage of temporary objects, cache objects, avoid not necessary allocation
- sync.Once == only first call
- sync.Cond == waiting for an event [Broadcast, Signal, Wait]
*/

type Dog struct {
	name         string
	walkDuration time.Duration
}

func (d Dog) Walk(wg *sync.WaitGroup) {
	fmt.Printf("%s is taking a walk\n", d.name)
	time.Sleep(d.walkDuration)
	fmt.Printf("%s is going home\n", d.name)
	wg.Done()
}

type httpPkg struct{}

func (httpPkg) Get(url string) {}

var http httpPkg

var i int // i == 0

func worker(wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	i = i + 1
	mu.Unlock()

	wg.Done()

}

func main() {
	dogs := []Dog{{"vasya", time.Second}, {"john", time.Second * 3}}
	runtime.GOMAXPROCS(4)
	var wg = &sync.WaitGroup{}
	var mu = &sync.Mutex{}
	for _, d := range dogs {
		wg.Add(1)
		go d.Walk(wg)
	}

	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			http.Get(url)
		}(url)
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go worker(wg, mu)
	}
	wg.Wait()

	fmt.Println("value of i after 1000 operations is", i)

	fmt.Println("everybody's home")

}
