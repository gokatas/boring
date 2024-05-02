// You can use channel as a handle on a service. Here we make Ann and Joe talk
// in lockstep, i.e. one after another.
package main

import (
	"fmt"

	"github.com/gokatas/boring"
)

func main() {
	ann := boring.Person("Ann")
	joe := boring.Person("Joe")
	for i := 0; i < 5; i++ {
		fmt.Println(<-ann)
		fmt.Println(<-joe)
	}
}
