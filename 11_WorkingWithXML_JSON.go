package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

type Person struct {
	Name string `json:"fullname,omitempty"`
	Age  int
	Job  struct {
		Department string
		Title      string
	}
}

/*
- empty interface -> map
- tags for json
*/

type Address struct {
	City, State string
}
type Person2 struct {
	XMLName   xml.Name `xml:"person"`
	Id        int      `xml:"id,attr"`
	FirstName string   `xml:"name>first"`
	LastName  string   `xml:"name>last"`
	Age       int      `xml:"age"`
	Height    float32  `xml:"height,omitempty"`
	Married   bool
	Address
	Comment string `xml:",comment"`
}

func main() {
	p1 := &Person{
		Name: "Vasya",
		Age:  36,
		Job: struct {
			Department string
			Title      string
		}{Department: "Operations", Title: "Boss"},
	}
	j, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("p1 json %s\n", j)
	var p2 Person
	json.Unmarshal(j, &p2)
	fmt.Printf("p2: %v\n", p2)

	j = []byte(`{"Name":"Vasya", "Job":{"Department":"Operations","Title":"Boss"}}`)
	var p3 interface{}
	json.Unmarshal(j, &p3)
	fmt.Printf("p3: %v\n", p3)
	person, ok := p3.(map[string]interface{})
	if ok {
		fmt.Printf("name=%s\n", person["Name"])
	}

	v := &Person2{Id: 13, FirstName: "John", LastName: "Doe", Age: 42}
	v.Comment = " Need more details. "
	v.Address = Address{"Hanga Roa", "Easter Island"}

	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("  ", "    ")
	enc.Encode(v)

}
