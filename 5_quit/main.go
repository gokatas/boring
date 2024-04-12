// Jack will quit talking when we let him now. He'll also let us know when he's
// done quitting.
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
