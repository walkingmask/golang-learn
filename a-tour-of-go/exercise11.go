package main

// https://go-tour-jp.appspot.com/concurrency/10
//http://qiita.com/peketamin/items/6a65cd9fec0205026afe#web-%E3%82%AF%E3%83%AD%E3%83%BC%E3%83%A9%E3%83%BC

import (
  "fmt"
  "sync"
)

type Fetcher interface {
  // Fetch returns the body of URL and
  // a slice of URLs found on that page.
  Fetch(url string) (body string, urls []string, err error)
}

type EUrl struct {
  urls map[string]int
  mux  sync.Mutex
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, eurl *EUrl) {
  // TODO: Fetch URLs in parallel.
  // TODO: Don't fetch the same URL twice.
  // This implementation doesn't do either:
  if depth <= 0 {
    return
  }
  // Check existence
  (*eurl).mux.Lock()
  if _, ok := (*eurl).urls[url]; ok {
    (*eurl).mux.Unlock()
    return
  }
  (*eurl).urls[url] = 1
  (*eurl).mux.Unlock()
  body, urls, err := fetcher.Fetch(url)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Printf("found: %s %q\n", url, body)
  for _, u := range urls {
    Crawl(u, depth-1, fetcher, eurl)
  }
  return
}

func main() {
  eurl := EUrl{urls: make(map[string]int)}
  Crawl("http://golang.org/", 4, fetcher, &eurl)
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
