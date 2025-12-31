// // Google I/O 2013 - Advanced Go Concurrency Patterns
package concurrency

import (
	"fmt"
	"time"
)

type Ball struct{ hits int }

func Play() {
	table := make(chan *Ball)
	go player("ping", table)
	go player("pong", table)

	// table <- new(Ball)
	time.Sleep(time.Second)
	<-table
}

func player(name string, table chan *Ball) {
	for {
		ball := <-table
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}
