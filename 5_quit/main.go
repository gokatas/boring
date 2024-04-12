package main

import (
	"fmt"
	"math/rand"

	"boring"
)

func main() {
	quit := make(chan bool)
	c := boring.Quitter("quiter", quit)
	for i := rand.Intn(10); i >= 0; i-- {
		fmt.Println(<-c)
	}
	quit <- true
	fmt.Println("quitting...")
	<-quit
}
