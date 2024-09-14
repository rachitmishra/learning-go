package coffeeshop

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func counter(orders chan *order) {
	for order := range orders {
		time.Sleep(10 * time.Millisecond)
		order.state = StateReceived
		orders <- order
	}
}

func grinder(orders chan *order) {
	for order := range orders {
		if order.state != StateReceived {
			continue
		}
		order.state = StateGrounding
		orders <- order
		time.Sleep(10 * time.Millisecond)
		order.state = StateGrounded
		orders <- order
	}
}

func coffeemachine(orders chan *order) {
	for order := range orders {
		if order.state != StateGrounded {
			continue
		}
		order.state = StateMaking
		orders <- order
		time.Sleep(10 * time.Millisecond)
		order.state = StateReady
		orders <- order
	}
}

func server(orders chan *order) {
	for order := range orders {
		if order.state != StateReady {
			continue
		}
		order.state = StateServing
		orders <- order
		time.Sleep(10 * time.Millisecond)
		order.state = StateServed
		orders <- order
	}
}

func display(orders <-chan *order) {
	for order := range orders {
		fmt.Printf("order state %s for order %d \n", order.state, order.id)
	}
}

func CoffeeShop() {
	orders := make(chan *order, 10)
	var wg sync.WaitGroup

	wg.Add(4)
	go counter(orders)
	go grinder(orders)
	// go coffeemachine(orders)
	// go server(orders)
	go display(orders)
	for {
		time.Sleep(1 * time.Second)
		orders <- newOrder(rand.Int32N(100))
	}
}
