package main

import (
	"fmt"

	"github.com/gen2brain/dlgs"
)

/*
- work stealing
- m:n == m goroutinr on n threads
- states: started, waited for start, not ready for start
- work stealing instead of work sharing
- global (1/61) -> local queue -> M -> P
- goroutine blocked: Gpsched(), read/write from network, system calls, mutexes/channels etc, morestack().
*/
func main() {
	item, _, err := dlgs.List("List", "Select item from list:", []string{"Bug", "New Feature", "Improvement"})
	if err != nil {
		panic(err)
	}
	fmt.Println(item)
}
