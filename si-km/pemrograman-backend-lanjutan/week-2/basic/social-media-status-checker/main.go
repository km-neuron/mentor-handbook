package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	websites := []string{
		"https://facebook.com",
		"https://twitter.com",
		"https://linkedin.com",
	}

	// check websites sequentially
	startTime := time.Now()
	checkWebsitesSeq(websites)
	elapsedTime := time.Since(startTime)
	fmt.Printf("Time taken sequentially: %s\n", elapsedTime)

	// check websites concurrently using goroutines
	startTime = time.Now()
	checkWebsitesConc(websites)
	elapsedTime = time.Since(startTime)
	fmt.Printf("Time taken concurrently (random): %s\n", elapsedTime)

	// check websites concurrently using channels
	startTime = time.Now()
	checkWebsitesChan(websites)
	elapsedTime = time.Since(startTime)
	fmt.Printf("Time taken concurrently (ordered): %s\n", elapsedTime)
}

func checkWebsitesSeq(websites []string) {
	for _, website := range websites {
		checkWebsite(website)
	}
}

func checkWebsitesConc(websites []string) {
	for _, website := range websites {
		go checkWebsite(website)
	}
	time.Sleep(2 * time.Second) // wait for all goroutines to finish
}

func checkWebsitesChan(websites []string) {
	results := make(chan string)

	for _, website := range websites {
		go func(website string) {
			status := checkWebsite(website)
			results <- status
		}(website)
	}

	for i := 0; i < len(websites); i++ {
		fmt.Println(<-results)
	}
}

func checkWebsite(website string) string {
	resp, err := http.Get(website)
	if err != nil {
		return fmt.Sprintf("[%s] %s is down", "ERR", website)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Sprintf("[%d] %s is down", resp.StatusCode, website)
	}
	return fmt.Sprintf("[%d] %s is up", resp.StatusCode, website)
}
