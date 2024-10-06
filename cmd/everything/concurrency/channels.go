package concurrency

import (
	"fmt"
)

func OnlyWrite(out chan<- int) {
	for i := 0; i < 5; i++ {
		out <- i
	}
}

func OnlyRead(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

func ChannelReadAndWrite() {
	out := make(chan int)
	in := make(chan int)
	go OnlyWrite(out)
	go OnlyRead(in)
	fmt.Println("All done!")
}

func ChannelRead() {
	// 1
	chan1 := make(chan string)
	go func() {
		chan1 <- "hello!"
	}()

	fmt.Println(<-chan1)
}

func RangeOverChannel() {
	// 2
	chan2 := make(chan string)
	go func() {
		chan2 <- "hello!"
		chan2 <- "world"
	}()

	for in := range chan2 {
		fmt.Println(in)
	}
}

func Channels() {
	RecieveFromClosedChannel()
}

func RecieveFromNilChannel() {
	var c chan string
	fmt.Println(<-c)
}

func SendToNilChannel() {
	var c chan string
	fmt.Println(<-c)
}

func SendToClosedChannel() {
	c := make(chan string, 10)
	close(c)
	c <- "hello"
}

// Closed channel never blocks
func RecieveFromClosedChannel() {
	c := make(chan string)
	close(c)
	say := <-c
	fmt.Println(say)
}
