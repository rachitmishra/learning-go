package concurrency

import (
	"fmt"
	"time"
)

func worker(data int) {
	fmt.Println(data)
}

func Goroutines() {
	// fork-join model
	go worker(1)
	go worker(2)
	go worker(3)
	time.Sleep(time.Second)
	fmt.Println("All done!")
}
