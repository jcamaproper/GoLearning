package main

import "fmt"

func main() {

	/* //card := "Ace of Spades"
	card := newCard()

	fmt.Println(card) */

	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	for _, card := range cards {

		fmt.Println(card)
	}

	fmt.Println(cards)
}

func newCard() string {

	return "Five of Diamonds"
}
