package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://truecaller.com",
	}

	ch := make(chan string)

	for _, link := range links {
		go checkLink(link, ch)
	}

	for l := range ch {
		go func(li string) {
			time.Sleep(1 * time.Second)
			checkLink(li, ch)
		}(l)
	}
}

func checkLink(link string, ch chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(time.Now(), link, "might be down!")
		ch <- link
		return
	}

	fmt.Println(link, "is up..")
	ch <- link
}
