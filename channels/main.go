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
	}

	c := make(chan string)

	for _, url := range links {
		go checkLink(url, c)
	}

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }
	// for {
	// 	go checkLink(<-c, c)
	// }
	for l := range c {
		go func(link string) {
			time.Sleep(time.Second * 5)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url + " might be down!")
		c <- url
		return
	}
	fmt.Println(url + " is working I think!")
	c <- url
}
