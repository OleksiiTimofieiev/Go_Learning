package main

import "fmt"

func Fast() int {
	acc := new(int)
	for i := 0; i < 10; i++ {
		acc2 := new(int)
		*acc2 = *acc + 1
		acc = acc2
	}
	return *acc
}

func Slow() int {
	acc := new(int)
	for i := 0; i < 1000; i++ {
		acc2 := new(int)
		*acc2 = *acc + 1
		acc = acc2
	}
	return *acc
}

func OnStack() {
	for i := 0; i < 1000; i++ {
		a := 100
		_ = a
	}
}

func main() {
	fmt.Println("--- Profiling ---")
}
