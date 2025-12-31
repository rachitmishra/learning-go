package handlers

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"

	o "rachitmishra.com/go/coffee-shop/data"
)

func counter(orders chan *o.Order) {
	select {
	case order := <-orders:
		fmt.Println("processing on counter")
		time.Sleep(10 * time.Millisecond)
		order.State = o.StateReceived
		fmt.Println("processing on processed")
		orders <- order
	default:
		fmt.Println("counter empty")
	}
}

func grinder(orders chan *o.Order) {
	select {
	case order := <-orders:
		fmt.Println("grinding on counter")
		time.Sleep(10 * time.Millisecond)
		order.State = o.StateGrounding
		orders <- order
		time.Sleep(10 * time.Millisecond)
		order.State = o.StateGrounded
		orders <- order
		fmt.Println("processing on grounded")
	default:
		fmt.Println("grinder empty")
	}
}

func coffeemachine(orders chan *o.Order) {
	select {
	case order := <-orders:
		fmt.Println("making on machine")
		if order.State != o.StateGrounded {
			return
		}
		order.State = o.StateMaking
		orders <- order
		time.Sleep(10 * time.Millisecond)
		order.State = o.StateReady
		orders <- order
		fmt.Println("making on ready")
	default:
		fmt.Println("machine empty")
	}
}

func server(orders chan *o.Order, wg *sync.WaitGroup) {
	select {
	case order := <-orders:
		fmt.Println("serving order")
		if order.State != o.StateReady {
			return
		}
		order.State = o.StateServing
		orders <- order
		time.Sleep(10 * time.Millisecond)
		order.State = o.StateServed
		orders <- order
		wg.Done()
	default:
		fmt.Println("server empty")
	}
}

func display(orders <-chan *o.Order) {
	select {
	case order := <-orders:
		fmt.Printf("order state %s for order %d \n", order.State, order.Id)
	default:
		fmt.Println("display empty")
	}
}

func CoffeeShop() {
	fmt.Println("welcome to coffee shop")
	orders := make(chan *o.Order)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go counter(orders)
	go grinder(orders)
	go coffeemachine(orders)
	go server(orders, &wg)
	go display(orders)
	go func() {
		for {
			wg.Add(1)
			fmt.Println("new order")
			orders <- o.NewOrder(rand.Int32N(100))
			time.Sleep(1 * time.Second)
		}
	}()
	wg.Wait()
}
