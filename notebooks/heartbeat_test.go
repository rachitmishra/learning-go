package main

import (
	"testing"
	"time"
)

func TestDoWork_GenerateInts(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	ints := []int{0, 1, 2, 3, 5}
	_, results := DoWork(done, ints...)

	for i, expected := range ints {
		select {
		case r:= <-results:
			if r != expected {
				t.Errorf("index %v: expected %v, but received %v,", i, expected, r)
			}
		case <-time.After(1 * time.Second):
			t.Fatal("test timed out")
		}
	}
}
