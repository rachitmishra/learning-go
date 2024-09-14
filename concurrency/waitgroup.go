package concurrency

import (
	"fmt"
	"sync"
)

func sayHello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello!")
}

func WaitGroup() {
	var wg sync.WaitGroup
	wg.Add(1)
	go sayHello(&wg)
	wg.Wait()
}

func WaitGroup1() {
	var wg sync.WaitGroup
	wg.Add(1)
	say := "hello"
	go func() {
		defer wg.Done()
		say = "welcome"
	}()
	wg.Wait()
	fmt.Println(say)
}

func WaitGroup2() {
	var wg sync.WaitGroup
	for _, say := range []string{"hello", "world", "go away"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(say)
		}()
	}
	wg.Wait()
}
