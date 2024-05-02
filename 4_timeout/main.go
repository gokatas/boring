// We use select to let talk Ann or Joe (whosever is ready) until no one speaks
// for 600ms.
//
// Select statement is another way to handle multiple channels. It's like switch
// but each case is a communication.
//
//   - All channels are evaluated.
//   - Blocks until one communication can proceed.
//   - If multiple can proceed, chooses (pseudo-)randomly.
//   - A default case, if present, executes immediately if no channel is ready.
package main

import (
	"fmt"
	"time"

	"github.com/gokatas/boring"
)

func main() {
	c := fanIn(boring.Person("Ann"), boring.Person("Joe"))
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-time.After(time.Millisecond * 600):
			fmt.Println("timeout")
			return
		}
	}
}

func fanIn(c1, c2 <-chan string) <-chan string {
	c := make(chan string)
	go func() { // only one goroutine is needed now
		for {
			select {
			case s := <-c1:
				c <- s
			case s := <-c2:
				c <- s
			}
		}
	}()
	return c
}
