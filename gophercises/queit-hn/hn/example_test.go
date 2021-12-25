package hn_test

import (
	"fmt"
	"qhn/hn"
)

func Example() {
	c := hn.NewClient("")
	ids, err := c.TopItems()
	if err != nil {
		panic(err)
	}
	for i := 0; i < 5; i++ {
		item, err := c.GetItem(ids[i])
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s (by %s)\n", item.Title, item.By)
	}
}
