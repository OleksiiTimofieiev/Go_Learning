package main

import (
	"fmt"
)

// int, uint, int8, uint8, ...
// byte uint8, rune = int32
// float32, float64
// complex64, ...
// string
// uintptr, *int, *string, ...

/*
** by default all is zero
** явное преобразование типов
** type Age uint32 == alias
** nil == NULL
 */

var Storage = make(map[string]string)
var storage = 3

type User struct {
	Name     string // public variable
	password string // private variable
}

func Answer() int {
	return 42
}

func main() {
	var i int = 10
	j := i

	/* strings */
	// strings.Builder
	s := "Hello"
	var c byte = s[0]
	var s2 string = s[0:4]
	s3 := s + "test"
	l := len(s3)

	/* unicod */
	var r rune = 'Я'

	fmt.Println("Hello, world!", j, s, c, s2, s3, l, r)

}
