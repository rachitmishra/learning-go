package concurrency

import "fmt"

func Channels() {
	var doneNil chan string
	// fmt.Println("Sending to nil channel ...")
	// doneNil <- "nil" // Will block
	fmt.Println("Reading from nil channel ...")
	<-doneNil // Will block
	var doneClosed = make(chan string)
	close(doneClosed)
	// fmt.Println("Sending to nil channel ...")
	// doneNil <- "nil" // Will block
	// fmt.Println("Reading from closed channel ...")
	// doneNil <- "nil" // Will block
	done := make(chan struct{})
	fmt.Println("Sending to channel ...")
	done <- struct{}{}
	// go func() {
	// 	fmt.Println("Sleeping for 2 second...")
	// 	time.Sleep(time.Second * 2)
	// 	fmt.Println("Done ...")
	// }()
	fmt.Println("Reading from channel ...")
	<-done
}
