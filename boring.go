/*
Boring contains various Go concurrency patterns.

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

const format = "%s: blah %d"

// Person is a generator, i.e. a function that returns a channel. A channel
// allows for communication and synchronization between goroutines. We launch
// the goroutine from inside the generator.
func Person(name string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf(format, name, i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
		}
	}()
	return c
}

// Quitter talks until we stop him by sending on the quit channel. He also
// notifies he's done quitting by sending back on the quit channel.
func Quitter(name string, quit chan bool) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			select {
			case c <- fmt.Sprintf(format, name, i):
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
