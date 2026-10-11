package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	sites := []string{
		"https://go.dev", "https://google.com", "https://wikipedia.org",
	}

	limit := make(chan struct{}, 3)
	var group sync.WaitGroup

	for _, site := range sites {
		group.Add(1)
		limit <- struct{}{}

		go func(link string) {
			defer group.Done()
			defer func() { <-limit }()

			resp, err := http.Get(link)
			if err != nil {
				fmt.Printf("%s -> Ошибка: %v\n", link, err)
				return
			}
			defer resp.Body.Close()

			fmt.Printf("%s -> %s\n", link, resp.Status)
		}(site)
	}
	group.Wait()
}
