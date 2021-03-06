package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type FetchedUrls struct {
	urls map[string]bool
	mux  sync.Mutex
}

func (f *FetchedUrls) Exists(url string) bool {
	return f.urls[url]
}

func (f *FetchedUrls) Add(url string) {
	f.mux.Lock()
	f.urls[url] = true
	f.mux.Unlock()
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, fetched *FetchedUrls, ret chan string) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	defer close(ret)

	if fetched.Exists(url) {
		return
	}

	if depth <= 0 {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		ret <- err.Error()
		fetched.Add(url)
		return
	}

	ret <- fmt.Sprintf("found: %s %q", url, body)
	fetched.Add(url)

	result := make([]chan string, len(urls))
	for i, u := range urls {
		result[i] = make(chan string)
		go Crawl(u, depth-1, fetcher, fetched, result[i])
	}

	for i := range result {
		if !fetched.Exists(urls[i]) {
			for s := range result[i] {
				ret <- s
			}
		}
	}
	return
}

func main() {
	result := make(chan string, 1)
	fetched := FetchedUrls{urls: make(map[string]bool)}
	go Crawl("http://golang.org/", 4, fetcher, &fetched, result)

	for s := range result {
		fmt.Println(s)
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
