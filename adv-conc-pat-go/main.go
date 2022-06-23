package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/SlyMarbo/rss"
)

// https://talks.golang.org/2013/advconc.slide#12

type Subscriber interface {
	Update() <-chan *rss.Item // stream of Items
	Close() error             // shuts down the stream
}

// sub implements the Subscriber interface.
type sub struct {
	*rss.Feed
	updates chan *rss.Item
	closing chan chan error
}

func (s *sub) Update() <-chan *rss.Item { return s.updates }

func (s *sub) Close() error {
	errc := make(chan error)
	s.closing <- errc
	return <-errc
}

// loop fetches items using s.Feed and sends them on s.updates. loop exits when
// s.Close is called.
func (s *sub) loop() {
	pending := make([]*rss.Item, len(s.Items))
	copy(pending, s.Items)
	next := s.Refresh
	var err error // set when Fetch fails
	var fetchDone chan fetchResult
	var fetchDelay time.Duration // initially 0 (no delay)
	for {
		var first *rss.Item
		var updates chan *rss.Item
		if len(pending) > 0 {
			first = pending[0]
			updates = s.updates
		}
		if now := time.Now(); next.After(now) {
			fetchDelay = next.Sub(now)
			fmt.Printf("fetchDelay %+v\n", fetchDelay)
		}
		var startFetch <-chan time.Time
		if fetchDone == nil {
			startFetch = time.After(fetchDelay)
			fmt.Printf("fetchDone startFetch %+v\n", startFetch)
		}

		select {
		case <-startFetch:
			fmt.Println("startFetch case")
			fetchDone = make(chan fetchResult, 1)
			go func() {
				err := s.Feed.Update()
				fetchDone <- fetchResult{fetched: s.Items, err: err}
			}()
		case result := <-fetchDone:
			fmt.Println("fetchDone case")
			fetchDone = nil
			next = s.Refresh
			fmt.Println("err", result.err)
			if result.err != nil {
				next = time.Now().Add(10 * time.Second)
			}
			fmt.Println("fetched", result.fetched)
			for _, i := range result.fetched {
				if _, seen := s.ItemMap[i.ID]; seen {
					fmt.Println("seen item before:", i.Title)
					continue
				}
				pending = append(pending, i)
			}
		case updates <- first:
			fmt.Println("updates case")
			pending = pending[1:]
		case errc := <-s.closing:
			fmt.Println("closing case")
			errc <- err
			close(s.updates) // tells receiver we're done
			return
		}
	}
}

type fetchResult struct {
	fetched []*rss.Item
	err     error
}

type feedURL string

// Subscribe fetches data and returns a Subscription that can update and close.
func Subscribe(furl feedURL) Subscriber {
	feed, err := rss.Fetch(string(furl))
	if err != nil {
		log.Fatal(err)
	}
	s := &sub{
		Feed:    feed,
		updates: make(chan *rss.Item),
		closing: make(chan chan error),
	}
	go s.loop()
	return s
}

// Merge aggregates multiple subscriptions into one subscription.
func Merge(subs ...Subscriber) Subscriber {
	return &mergeSub{subs: subs, errc: make(chan error)}
}

type mergeSub struct {
	subs []Subscriber
	errc chan error
}

func (ms *mergeSub) Update() <-chan *rss.Item {
	updates := make(chan *rss.Item)
	for _, s := range ms.subs {
		s := s
		go func() {
			log.Println("mistakes")
			for it := range s.Update() {
				updates <- it
			}
		}()
	}
	return updates
}

func (ms *mergeSub) Close() error {
	for _, s := range ms.subs {
		if err := s.Close(); err != nil {
			ms.errc <- err
		}
	}
	if err := <-ms.errc; err != nil {
		return err
	}
	return nil
}

func main() {
	// FIXME(jay): Merge sucks
	s := Merge(Subscribe("https://xkcd.com/atom.xml"))
	for it := range s.Update() {
		fmt.Println(it)
	}
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt)
	// feed, err := rss.Fetch("https://xkcd.com/atom.xml")
	// if err != nil {
	// 	panic(err)
	// }
	// for _, it := range feed.Items {
	// 	fmt.Println(it)
	// }
	<-exit
	if err := s.Close(); err != nil {
		panic(err)
	}
}
