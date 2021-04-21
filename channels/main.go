package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://www.google.com",
		 "http://www.amazon.co.uk", 
		 "http://www.facebook.com",
		 "http://www.stackoverflow.com",
		 "http://www.golang.org",
		}

	c := make(chan string)

	for _, website := range links {
		go checkWebsite(website, c)
	}

	for l := range c {
		go func() {
			time.Sleep(5 * time.Second)
			go checkWebsite(l, c)
		}()
	}
}

func checkWebsite(website string, c chan string) {
	_, err := http.Get(website)
	if err != nil {
		fmt.Println(website, "our website might be down...")
		c  <- website
		return
	}
	fmt.Println(website, "is working fine...")
	c  <- website
}