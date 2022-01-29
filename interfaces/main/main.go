package main

import "fmt"

type bot interface {
	getGreeting() string
	printVersion() float64
}

type englishBot struct{}
type spanishBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
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
