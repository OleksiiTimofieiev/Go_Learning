//go:generate echo "hello world"
//go:generate ./command.sh
//go:generate ls -la
//go:generate -command bye echo "Bye"
//go:generate bye

package main

/*
- go generate -v -x -n -run enums (regexp)
- generate structures from JSON
- generate mocks for tests
- Protobug: generate code from protocol description: .proto
- bindata: adding binary data (pictures or else) in go code in
form of []byte
- пщ
*/

import (
	"fmt"
)

/*
- go generate
*/
func main() {
	fmt.Println("--- Code Generation ---")
}
