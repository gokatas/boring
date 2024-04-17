// The main goroutine communicates with the goroutine launched by Person via a
// channel. A channel allows for communication and synchronization between
// goroutines.
package main

import (
	"boring"
	"fmt"
)

func main() {
	c := boring.Person("Joe")
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
}
