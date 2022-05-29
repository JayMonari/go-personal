package main

import (
	"fmt"
	"time"

	"github.com/SlyMarbo/rss"
)

// https://talks.golang.org/2013/advconc.slide#12

func main() {
	feed, err := rss.Fetch("https://go.dev/blog/feed.atom")
	if err != nil {
		panic(err)
	}
	fmt.Println(feed)
	time.Sleep(10 * time.Second)
}
