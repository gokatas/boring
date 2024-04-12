// Lockstep makes Ann and Joe talk in lockstep (one after another). You can use
// channels as a handle on a service.
package main

import (
	"fmt"

	"boring"
)

func main() {
	ann := boring.Person("Ann")
	joe := boring.Person("Joe")
	for i := 0; i < 5; i++ {
		fmt.Println(<-ann)
		fmt.Println(<-joe)
	}
}
