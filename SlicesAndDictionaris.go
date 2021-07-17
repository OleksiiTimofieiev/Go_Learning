package main

import (
	"fmt"
)

func main() {
	/* array */

	var arr1 [3]int
	arr1[0] = 1

	arr2 := [10]int{4, 4, 4}

	fmt.Print("Arrays", arr1, arr2)

	/* slice (kinda array with realloc when needed == maybe kinda vector) */
	/* during copy -> same address, if more then capacity -> realloc of memory */
	var s []int

	s = append(s, 1)
	s = append(s, 2, 3)

	fmt.Println("\nSlices", s)
	/* sort */
	/* to sort the Interface has to be implemented */

	/* map == hash map*/
	cache := map[string]string{
		"one": "one",
		"two": "two",
	}

	value, ok := cache["one"]

	/* порядок ключей рандомизирован */
	for key, val := range cache {
		fmt.Println(key, val)
	}

	delete(cache, "two")

	/* get keys or values */
	/* go don`t have generics */
	/* less malloc -> quicker the program */

	/* key == any data type with == operator defined */

	values := make([]string, 0, len(cache))

	for _, val := range cache {
		values = append(values, val)
	}

	fmt.Println(cache, value, ok, values)

}
