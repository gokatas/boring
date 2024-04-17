// Quit makes a boring person stop talking by sending on a channel. They will
// let us know when they're done quitting via another channel.
package main

import (
	"fmt"
	"math/rand"

	"boring"
)

func main() {
	quit := make(chan bool)
	c := boring.Quitter("Jack", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- true
	fmt.Println("quitting...")
	<-quit
}
