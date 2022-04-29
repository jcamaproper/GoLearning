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
		"https://proper.ai",
		"https://www.linkedin.com/feed"}

	// On this way, the code does no wait each link to complete, run each link in parallel

	for _, link := range links {
		go checkLink(link, c)

	}
	// wait for a value to be sent into the channel. When we get one, log it out inmediately
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " ", "Its down :c")
		// Send value to the channel
		c <- "Down"
		return
	}
	fmt.Println(link, " ", "Its up! =D")
	// Send value to the channel
	c <- "Up"
}
