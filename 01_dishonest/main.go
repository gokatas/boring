// This is the first boring conversation. It's not an honest one because there
// is no communication between the main and the say goroutine.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	go say("blah")
	time.Sleep(time.Second * 5)
}

func say(msg string) {
	for i := 0; ; i++ {
		fmt.Printf("%s, %d\n", msg, i)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
	}
}
