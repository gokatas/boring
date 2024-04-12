// The main goroutine communicates with a goroutine started by Person.
package main

import (
	"fmt"

	"boring"
)

func main() {
	c := boring.Person("Joe")
	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}
}
