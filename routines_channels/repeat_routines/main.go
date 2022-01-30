package main

import (
	"fmt"
	"net/http"
)

func main() {

	// channel to comunicate throught go routines
	c := make(chan string)
	links := []string{
		"https://www.record.com.mx",
		"https://www.facebook.com",
		"https://www.stackoverflow.com",
		"https://www.golang.org",
		"https://www.amazon.com",
		"https://www.linkedin.com/feed"}

	// On this way, the code does no wait each link to complete, run each link in parallel

	for _, link := range links {
		go checkLink(link, c)

	}
	// infinite loop still bloking channel
	for l := range c {
		//go checkLink(<-c, c)
		go checkLink(l, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " ", "Its down")
		// Send value to the channel
		c <- link
		return
	}
	fmt.Println(link, " ", "Its up")
	// Send value to the channel
	c <- link
}
