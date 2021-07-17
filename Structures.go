package main

import (
	"fmt"
)

type Signal struct{}

type User struct {
	Id       int64
	Name     string
	previous *User
	counter  *int
	test     struct {
		test string
	}
}

var counter int

func NewUser() *User {
	counter++
	return &User{counter: &counter}
}

func main() {

	var John User
	var Ivan *User = NewUser()
	Ivan.counter = &counter
	user3 := NewUser()
	user4 := User{123, "asdf", user3, &counter, struct{ test string }{test: "sdf"}}

	John.Id = 123
	John.Name = "awd"

	/* anonymous structure */
	var wordCounts []struct {
		w string
		i int
	}

	wordCounts = append(wordCounts, struct {
		w string
		i int
	}{w: "qwe", i: 4})

	wordCounts = append(wordCounts, struct {
		w string
		i int
	}{w: "sdf", i: 4})

	fmt.Println("Structures", John)
	fmt.Println("Structures", Ivan)

	fmt.Println("Structures", user3)
	fmt.Println("Structures", user4)
	fmt.Println("Structures", wordCounts)

}
