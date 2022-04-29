package main

import (
	"fmt"
	"net/http"
	"os"
)

// interface can be inside another interface
/* type masterBot interface {
	bot
} */

/* type bot interface {
	getGreeting() string
	printVersion() float64
}

type englishBot struct{}
type spanishBot struct{} */

func main() {

	/* 	eb := englishBot{}
	   	sb := spanishBot{}

	   	printGreeting(eb)
	   	printGreeting(sb) */

	resp, err := http.Get("https://coinmarketcap.com/")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	bs := make([]byte, 99999)

	resp.Body.Read(bs)
	fmt.Println(string(bs))

}

/* func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {

	return "Hi There!"
}

func (spanishBot) printVersion() float64 {
	return 1.0
}

func (englishBot) printVersion() float64 {
	return 1.2
}

func (spanishBot) getGreeting() string {

	return "Hola!"
}
*/
