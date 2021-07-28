package main

import (
	"fmt"
)

/* 57:45 */

/*
- multiple go routines have own small stack allocated on heap
- gc == memory allocation on arena
- SP == Stack Pointer in Go
- Arena = 64 Mb on linux -> Pages == 8kb -> Span == set of pages === Arena [ Span [ Pages ] ]
- Span: little/big, scan/noscan, for a specific size of the object
- nheap [ mcentral: free or full spans, mTreap, mCache for each CPU core ]
- span found -> check the bitmap to check free cells
- mcache -> mcentral -> mTreap
- GC: mark & sweep => fully stops everithing
- GC: three color mark (black, grey, white) & sweep => not stops everithing

*/

func main() {
	fmt.Println("test")
}
