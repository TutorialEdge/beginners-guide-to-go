package main

import (
	"fmt"
	"log"
	"net/http"

	"sync/atomic"
)

var processed int64

type Site struct {
	URL string
}

type Result struct {
	URL    string
	Status int
}

func crawl(wId int, jobs <-chan Site, results chan<- Result) {
	for site := range jobs {
		log.Printf("Worker ID: %d\n", wId)
		resp, err := http.Get(site.URL)
		if err != nil {
			log.Println(err.Error())
		}
		atomic.AddInt64(&processed, 1)
		results <- Result{
			URL:    site.URL,
			Status: resp.StatusCode,
		}
	}
}

func main() {
	fmt.Println("Atomic Counters in Go")

	jobs := make(chan Site, 3)
	results := make(chan Result, 3)

	for w := 1; w <= 3; w++ {
		go crawl(w, jobs, results)
	}

	urls := []string{
		"https://tutorialedge.net",
		"https://tutorialedge.net/pricing/",
		"https://example.com",
		"https://google.com",
	}

	for _, url := range urls {
		jobs <- Site{URL: url}
	}
	close(jobs)

	for a := 1; a <= 4; a++ {
		result := <-results
		log.Println(result)
	}

	fmt.Printf("URLs Processed: %d\n", processed)

}
