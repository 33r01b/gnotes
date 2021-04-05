package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	var urls = []string{
		"http://ozon.ru",
		"https://ozon.ru",
		"http://google.com",
		"http://somesite.com",
		"http://non-existent.domain.tld",
		"https://ya.ru",
		"http://ya.ru",
		"http://ёёёё",
	}

	client := &http.Client{Timeout: time.Second * 3}
	results := make(chan string)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	goroutines := make(chan bool, 5)

	go func() {
		var wg sync.WaitGroup
		for _, url := range urls {
			goroutines <- true
			wg.Add(1)
			go func(url string) {
				results <- parse(ctx, client, url)
				<-goroutines
				wg.Done()
			}(url)
		}
		wg.Wait()
		close(results)
	}()

	i := 1
	for r := range results {
		fmt.Println(r)

		i++
		if i > 2 {
			cancel()
			break
		}
	}

	fmt.Println("stopped")
}

func parse(ctx context.Context, client *http.Client, url string) string {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("%v", err)
	}

	req = req.WithContext(ctx)
	resp, err := client.Do(req)

	if err != nil {
		return fmt.Sprintf("Error %s\n", url)
	}

	return fmt.Sprintf("%d %s\n", resp.StatusCode, url)
}
