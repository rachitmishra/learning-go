package concurrency

import (
	"fmt"
	"time"

	"github.com/mmcdole/gofeed"
)

type Fetcher interface {
	Fetch() (feed *gofeed.Feed, err error)
}

type fetcher struct {
	parser *gofeed.Parser
	domain string
}

func (f *fetcher) Fetch() (feed *gofeed.Feed, err error) {
	feed, err = f.parser.ParseURL(f.domain)
	return feed, err
}

type Subscription interface {
	Updates() <-chan *gofeed.Item
	Close() error
}

type sub struct {
	fetcher Fetcher
	updates chan *gofeed.Item
	closed  bool
	err     error
}

func (s *sub) Updates() <-chan *gofeed.Item {
	return s.updates
}

func (s *sub) Close() error {
	close(s.updates)
	return nil
}

func Subscribe(fetcher Fetcher) Subscription {
	s := &sub{
		fetcher: fetcher,
		updates: make(chan *gofeed.Item),
	}
	go s.loop()
	return s
}

func (s *sub) loop() {
	if s.closed {
		close(s.updates)
		return
	}
	feed, err := s.fetcher.Fetch()
	if err != nil {
		s.err = err
		time.Sleep(10 * time.Second)
	}
	for _, item := range feed.Items {
		s.updates <- item
	}
	// if now := time.Now(); next.After(now) {
	// 	time.Sleep(next.Sub(now))
	// }
}

type merge struct {
	subs    []Subscription
	updates chan *gofeed.Item
}

func Merge(subs ...Subscription) Subscription {
	m := &merge{
		subs:    subs,
		updates: make(chan *gofeed.Item),
	}
	for _, sub := range subs {
		go func(s Subscription) {
			for it := range s.Updates() {
				m.updates <- it
			}
		}(sub)
	}
	return m
}

func (m *merge) Updates() <-chan *gofeed.Item {
	return m.updates
}

func (m *merge) Close() (err error) {
	for _, sub := range m.subs {
		if e := sub.Close(); err == nil && e != nil {
			err = e
		}
	}
	close(m.updates)
	return
}

func FetchRssFeed() {
	fetcher := &fetcher{
		parser: gofeed.NewParser(),
	}
	merged := Merge(
		Subscribe(fetcher),
		Subscribe(fetcher),
	)

	time.AfterFunc(3*time.Second, func() {
		fmt.Println("closed:", merged.Close())
	})

	for it := range merged.Updates() {
		fmt.Println(it.Link, it.Title)
	}

	panic("stacks")
}
