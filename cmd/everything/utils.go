package main

import (
	"errors"
	"fmt"
	"time"
)

func Retry(cmd func() error, maxRetry int, startBackOff, maxBackOff time.Duration) {
	for attempt := 0; ; attempt++ {
		if err := cmd(); err == nil {
			return
		}
		if attempt == maxRetry-1 {
			return
		}
		fmt.Printf("REtrying after %s\n", startBackOff)
		time.Sleep(startBackOff)
		if startBackOff < maxBackOff {
			startBackOff *= 2
		}
	}
}

func RunTask() {
	fn := func(a, b int) (int, error) {
		fmt.Printf("Function called with a: %d and b: %d\n", a, b)
		return 42, errors.New("some error")
	}
	var res int
	var err error
	Retry(func() error {
		res, err = fn(42, 100)
		return err
	},
		3,
		1*time.Second,
		4*time.Second)
	fmt.Println(res, err)
}
