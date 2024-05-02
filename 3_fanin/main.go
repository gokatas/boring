// FanIn function combines (multiplexes) two channels into one,
// letting talk Ann or Joe, whosever is ready.
package main

import (
	"fmt"

	"github.com/gokatas/boring"
)

func main() {
	c := fanIn(boring.Person("Ann"), boring.Person("Joe"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
}

func fanIn(c1, c2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-c1
		}
	}()
	go func() {
		for {
			c <- <-c2
		}
	}()
	return c
}
