// Quit makes a person stop talking by sending on a channel. They will
// also let us know when they're done quitting via another channel.
package main

import (
	"fmt"
	"math/rand"
	"time"

	"boring"
)

func main() {
	quit := make(chan bool)
	c := quitter("Jack", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- true
	fmt.Println("quitting...")
	<-quit
}

func quitter(name string, quit chan bool) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf(boring.Format, name, i):
			case <-quit:
				cleanup()
				quit <- true
				return
			}
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
		}
	}()
	return c
}

func cleanup() { time.Sleep(time.Millisecond * time.Duration(rand.Intn(3e3))) }
