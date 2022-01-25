package main

import "fmt"

// Create a new type of deck
// which is a slice of string
type deck []string

func main() {

	//cards := []string{"Ace of Diamonds", newCard()}
	//cards = append(cards, "Six of Spades")
	cards := deck{"Ace of Diamonds", newCard()}
	/* for _, card := range cards {

		fmt.Println(card)
	} */

	cards.print()
	fmt.Println(cards)
}

func newCard() string {

	return "Five of Diamonds"
}

func (d deck) print() {

	for i, card := range d {

		fmt.Println(i, card)
	}

}
