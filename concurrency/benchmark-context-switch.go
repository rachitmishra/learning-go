package concurrency

import (
	"sync"
	"testing"
)

func BenchmarkContextSwitch(b *testing.B) {
	var wg sync.WaitGroup
	begin := make(chan interface{})
	c := make(chan interface{})
	var token interface{}
	sender := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			c <- token
		}
	}
	receiver := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			<-c
		}
	}
	wg.Add(2)
	go sender()
	go receiver()
	b.StartTimer()
	close(begin)
	wg.Wait()
}
