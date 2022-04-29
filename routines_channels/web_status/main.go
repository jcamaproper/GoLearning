package main

import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{
		"https://www.record.com.mx",
		"https://www.facebook.com",
		"https://www.stackoverflow.com",
		"https://www.golang.org",
		"https://www.amazon.com"}

	// In this way, each request is done one by one, SERIAL Request
	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " ", "Its down :c")
		return
	}
	fmt.Println(link, " ", "It up! =D")
}
