package main

import (
	"log"
	"net/http"
)

type Site struct {
	URL string
}

type Result struct {
	Status int
}

func crawl(id int, jobs <-chan Site, results chan<- Result) {
	for site := range jobs {
		log.Printf("Worker ID: %d\n", id)
		resp, err := http.Get(site.URL)
		if err != nil {
			log.Println(err.Error())
		}
		results <- Result{Status: resp.StatusCode}
	}
}

func main() {
	jobs := make(chan Site, 3)
	results := make(chan Result, 3)

	for w := 1; w <= 3; w++ {
		go crawl(w, jobs, results)
	}

	urls := []string{
		"https://tutorialedge.net",
		"https://tutorialedge.net/pricing/",
		"https://example.com/",
		"https://google.com/",
	}

	for _, url := range urls {
		jobs <- Site{URL: url}
	}
	close(jobs)

	for a := 1; a <= 4; a++ {
		result := <-results
		log.Println(result)
	}
}
