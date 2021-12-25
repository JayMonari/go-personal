package main

import (
	"errors"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"qhn/hn"
	"sort"
	"strings"
	"sync"
	"time"
)

func main() {
	var port, numStories int
	flag.IntVar(&port, "port", 3000, "the port to start the web server on")
	flag.IntVar(&numStories, "amount", 30, "the number of top stories to display")
	flag.Parse()

	tpl := template.Must(template.ParseFiles("./index.html"))
	http.HandleFunc("/", handler(numStories, tpl))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func handler(numStories int, tpl *template.Template) http.HandlerFunc {
	sc := storyCache{
		amtStories: numStories,
		d:          3 * time.Minute,
	}
	go func() {
		ticker := time.NewTicker(3 * time.Second)
		for {
			temp := storyCache{
				amtStories: numStories,
				d:          3 * time.Minute,
			}
			temp.stories()
			sc.mu.Lock()
			sc.cache = temp.cache
			sc.expiration = temp.expiration
			sc.mu.Unlock()
			<-ticker.C
		}
	}()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		stories, err := sc.stories()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := templateData{
			Stories: stories,
			Time:    time.Now().Sub(start),
		}
		err = tpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Failed to process the template", http.StatusInternalServerError)
			return
		}
	})
}

type storyCache struct {
	amtStories int
	cache      []item
	expiration time.Time
	d          time.Duration
	mu         sync.Mutex
}

func (sc *storyCache) stories() ([]item, error) {
	sc.mu.Lock()
	defer sc.mu.Unlock()

	if time.Now().Sub(sc.expiration) < 0 {
		return sc.cache, nil
	}

	stories, err := getTopStoriesN(sc.amtStories)
	if err != nil {
		return nil, err
	}
	sc.expiration = time.Now().Add(sc.d)
	sc.cache = stories
	return sc.cache, nil
}

// getTopStoriesN returns n amount of stories from Hacker News API
func getTopStoriesN(n int) (stories []item, err error) {
	c := hn.NewClient("")
	ids, err := c.TopItems()
	if err != nil {
		return nil, errors.New("Failed to load top stories")
	}

	at := 0
	for len(stories) < n {
		need := (n - len(stories)) * 5 / 4
		stories = append(stories, getStories(ids[at:at+need])...)
		at += need
	}
	return stories, nil
}

func getStories(ids []int) []item {
	stories := make([]item, 0, 30)
	c := hn.NewClient("")
	type result struct {
		item item
		err  error
		idx  int
	}
	resCh := make(chan result)
	for i := 0; i < len(ids); i++ {
		go func(i, id int) {
			hnItem, err := c.GetItem(id)
			if err != nil {
				resCh <- result{idx: i, err: err}
			}
			resCh <- result{idx: i, item: parseHNItem(hnItem)}
		}(i, ids[i])
	}

	rr := make([]result, len(ids))
	for i := 0; i < len(ids); i++ {
		rr[i] = <-resCh
	}
	sort.Slice(rr, func(i, j int) bool {
		return rr[i].idx < rr[j].idx
	})

	for _, r := range rr {
		if r.err != nil {
			continue
		}
		if r.item.Type == "story" && r.item.URL != "" {
			stories = append(stories, r.item)
		}
		if len(stories) == cap(stories) {
			break
		}
	}
	return stories
}

func parseHNItem(it hn.Item) item {
	ret := item{Item: it}
	if pURL, err := url.Parse(ret.URL); err == nil {
		ret.Host = strings.TrimPrefix(pURL.Hostname(), "www.")
	}
	return ret
}

// item adds the Host field to hn.Item
type item struct {
	hn.Item
	Host string
}

type templateData struct {
	Stories []item
	Time    time.Duration
}
