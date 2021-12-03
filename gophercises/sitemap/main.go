package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"sitemap/link"
	"strings"
)

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

type loc struct {
	Value string `xml:"loc"`
}

type urlset struct {
	URLs  []loc  `xml:"url"`
	XMLns string `xml:"xmlns,attr"`
}

func main() {
	flagURL := flag.String("url", "https://gophercises.com", "the url that you want to build a sitemap for")
	maxDepth := flag.Int("depth", 3, "the maximum clicks to make on each page")
	flag.Parse()

	hrefs := bfs(*flagURL, *maxDepth)
	toXML := urlset{URLs: make([]loc, len(hrefs)), XMLns: xmlns}

	for i, h := range hrefs {
		toXML.URLs[i] = loc{h}
	}

	fmt.Print(xml.Header)
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "  ")
	if err := enc.Encode(toXML); err != nil {
		panic(err)
	}
}

type set map[string]struct{}

// https://dave.cheney.net/2014/03/25/the-empty-struct

func bfs(urlStr string, maxDepth int) []string {
	seen := make(set)
	var q set
	nq := set{urlStr: struct{}{}}
	for i := 0; i < maxDepth; i++ {
		q, nq = nq, make(map[string]struct{})
		if len(q) == 0 {
			break
		}
		for url := range q {
			if _, ok := seen[url]; ok {
				continue
			}
			seen[url] = struct{}{}

			for _, link := range getLinks(url) {
				if _, ok := seen[link]; !ok {
					nq[link] = struct{}{}
				}
			}
		}
	}

	links := make([]string, 0, len(seen))
	for lnk := range seen {
		links = append(links, lnk)
	}
	return links
}

func getLinks(uri string) []string {
	resp, err := http.Get(uri)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	reqURL := resp.Request.URL
	baseURL := &url.URL{
		Scheme: reqURL.Scheme,
		Host:   reqURL.Host,
	}

	return filter(parseHRefs(resp.Body, baseURL.String()), withPrefix(baseURL.String()))
}

func parseHRefs(r io.Reader, base string) []string {
	var hrefs []string
	links, _ := link.Parse(r)
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.HRef, "/"):
			hrefs = append(hrefs, base+l.HRef)
		case strings.HasPrefix(l.HRef, "http"):
			hrefs = append(hrefs, l.HRef)
		}
	}
	return hrefs
}

func filter(links []string, keepFn func(string) bool) []string {
	var ret []string
	for _, link := range links {
		if keepFn(link) {
			ret = append(ret, link)
		}
	}
	return ret
}

func withPrefix(p string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, p)
	}
}
