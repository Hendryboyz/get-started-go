package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string, r chan FetchResult) // (body string, urls []string, err error)
}

type UrlRepository struct {
	mu sync.Mutex
	urls map[string]bool
}

func (r *UrlRepository) IsVisited(url string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	_, existed := r.urls[url]
	return existed
}

func (r *UrlRepository) Add(url string) {
	r.mu.Lock()
	r.urls[url] = true
	r.mu.Unlock()
}

var urlRepo UrlRepository = UrlRepository{urls: make(map[string]bool)}

type FetchResult struct {
	body string
	urls []string
	err error
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// This implementation doesn't do either:
	if depth <= 0 {
		// Prevent from fetching the same URL twice.
		return
	}
	r := make(chan FetchResult)
	go fetcher.Fetch(url, r)
	result := <- r
	body, urls, err := result.body, result.urls, result.err
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {	
		Crawl(u, depth-1, fetcher)
	}
	return
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func createFetchResult(body string, urls []string, err error) FetchResult {
	return FetchResult{
		body: body,
		urls: urls,
		err: err,
	}
}

func (f fakeFetcher) Fetch(url string, r chan FetchResult) {
	if isVisited := urlRepo.IsVisited(url); isVisited {
		r <- createFetchResult("", nil, fmt.Errorf(""))
		return
	}
	urlRepo.Add(url)
	if res, ok := f[url]; ok {
		r <- createFetchResult(res.body, res.urls, nil)
		return
	}
	r <- createFetchResult("", nil, fmt.Errorf("not found: %s", url))
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
