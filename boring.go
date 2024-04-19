/*
Package boring contains basic Go concurrency patterns.

It's based on a talk by Rob Pike (2012):

  - video: https://www.youtube.com/watch?v=f6kdp27TYZs
  - slides: https://talks.golang.org/2012/concurrency.slide
  - code: https://talks.golang.org/2012/concurrency/support
*/
package boring

import (
	"fmt"
	"math/rand"
	"time"
)

const Format = "%s: blah %d"

// Person is a (bullshit :-) generator, i.e. a function that returns a channel.
// Person launches a goroutine that sends on the channel.
func Person(name string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf(Format, name, i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
		}
	}()
	return c
}
