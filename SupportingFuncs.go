package main

import (
	"fmt"
)

/* closure */
func Logger(prefix string) func(s string) {
	return func(s string) {
		fmt.Printf("[%s] %s\n", prefix, s)
	}
}
