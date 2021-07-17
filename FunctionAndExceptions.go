package main

import (
	"errors"
	"fmt"
)

func A() {
	defer fmt.Println("A")
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic: %+v\n", r)
		}
	}()
	B()
}
func B() {
	defer fmt.Println("B")
	C()
}
func C() {
	defer fmt.Println("C")
	Break()
}
func Break() {
	defer fmt.Println("D")
	panic(errors.New("the show must go on"))
}

/* closure */

func logger(prefix string) func(s string) {
	i := 0 /* kinda static variable as a result */
	return func(s string) {
		i++
		fmt.Printf("[%s] %s %d\n", prefix, s, i)
	}
}

/* methods */

type Employee struct {
	name, surname string
}

func Fullname(e Employee) string {
	return e.name + " " + e.surname
}

func (e Employee) Fullname() string {
	return e.name + " " + e.surname
}

func main() {
	defer fmt.Println("defer test")

	warn := logger("WARN")
	err := logger("ERROR")

	warn("Test")
	err("Err")
	err("Err")
	err("Err")

	println(Employee{"Al", "Pachino"}.Fullname())
	println(Fullname(Employee{"Al", "Pachino"}))

	/* defer (LIFO) */
	/* panic */
	/* recover */
	/* try catch analogue */
	A()
}
