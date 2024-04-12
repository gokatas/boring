package main

import (
	"fmt"

	"boring"
)

func main() { // the main goroutine
	c := boring.Person("Joe")
	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}
}
