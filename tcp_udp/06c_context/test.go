package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Millisecond)

	wg.Add(1)

	go dealLongWithCtx(wg, ctx)
	// cancel() == stop channel
	wg.Wait()
}

func dealLongWithCtx(wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randTimer := time.Duration(r.Intn(5000)) * time.Microsecond
	fmt.Println(randTimer)
	timer := time.NewTimer(randTimer)

	select {
	case <-ctx.Done():
		fmt.Println("Rejected")
	case <-timer.C:
		fmt.Println("Done")
	}
}
