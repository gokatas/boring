// The main goroutine communicates with a goroutine launched by Person via a
// channel.
package main

import (
	"boring"
	"fmt"
)

func main() {
	c := boring.Person("Joe")
	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}
}
