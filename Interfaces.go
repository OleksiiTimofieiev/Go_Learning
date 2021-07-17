package main

import (
	"fmt"
)

type shape interface {
	Area() float64
	Perimeter() float64
}

type Duck interface {
	Talk() string
	Walk()
	Swim()
}

type Dog struct {
	name string
}

/* all what Duck can is being a Duck */
func (d Dog) Talk() string {
	return "ARRRRRRR"
}

func (d Dog) Walk() {}
func (d Dog) Swim() {}

func quack(d Duck) {
	fmt.Println(d.Talk())
}

/* неявные интерфейсы  */
type MyStringer struct{ s string }

func (s MyStringer) String() string {
	return "MyStringer" + s.s
}

type Stringer interface {
	String() string
}

/* type can implement several interfaces */
type Hound interface {
	Hunt()
}

type Poodle interface {
	Bark()
}

type GoldenRetriever struct{ name string }

func (GoldenRetriever) Hunt() { fmt.Println("hunt") }
func (GoldenRetriever) Bark() { fmt.Println("bark") }

func f1(i Hound)  { i.Hunt() }
func f2(i Poodle) { i.Bark() }

/* one interface can have several types */
type Scandinav struct{ name string }

func (Scandinav) Bark() { fmt.Println("bark" + "-sc-" + "name") }

type ToyPoodle struct{ name string }

func (ToyPoodle) Bark() { fmt.Println("bark" + "-toy-" + "name") }

/* composition */
type Greeter interface {
	Hello()
}

type Stranger interface {
	Bye() string
	Greeter
	fmt.Stringer
}

/* empty interface */
// interface{}
func PrintAll(vals []interface{}) {
	for _, val := range vals {
		fmt.Println(val)
	}
}

func main() {
	log := Logger("Interfaces")
	log("Interfaces tests start")
	fmt.Println("Interfaces")
	quack(Dog{})
	fmt.Println(MyStringer{"Hello"})

	t := GoldenRetriever{"Jack"}

	f1(t)
	f2(t)

	var sc, toy Poodle

	sc = Scandinav{"Tom"}
	toy = ToyPoodle{"Russ"}

	sc.Bark()
	toy.Bark()

	log("Empty interface - kinda generic")

	vals := []interface{}{"John", "Pepe", "Huan"}
	// /* need to be converted to the interface type */
	names := []string{"asdf", "fda"}
	vals2 := make([]interface{}, len(names))
	for i, v := range names {
		vals2[i] = v
	}
	PrintAll(vals)
	PrintAll(vals2)

	defer log("Interfaces tests end")

}
