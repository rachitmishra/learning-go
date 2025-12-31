package concurrency

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
)

type counter struct {
	work chan int
	res  chan int
	eg   *errgroup.Group
}

func newCounter(ctx context.Context) *counter {
	eg, _ := errgroup.WithContext(ctx)
	c := &counter{
		work: make(chan int),
		res:  make(chan int),
		eg:   eg,
	}
	go c.recv(ctx)
	return c
}

func (c *counter) recv(ctx context.Context) error {
	sum := 0
	for {
		select {
		case i, ok := <-c.work:
			if !ok {
				c.res <- sum
				return nil
			}
			sum += i
		case <-ctx.Done():
			return context.Canceled
		}
	}
}

func (c *counter) send(ctx context.Context, n int) {
	c.eg.Go(func() error {
		for i := 0; i < n; i++ {
			select {
			case c.work <- i:
			case <-ctx.Done():
				return context.Canceled
			}
		}
		return nil
	})
}

func (c *counter) finalize() (int, error) {
	if err := c.eg.Wait(); err != nil {
		return 0, err
	}
	close(c.work)
	return <-c.res, nil
}

func Main() {
	ctx := context.Background()
	c := newCounter(ctx)
	c.send(ctx, 10)
	c.send(ctx, 10)
	c.send(ctx, 10)
	sum, err := c.finalize()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sum)
}
