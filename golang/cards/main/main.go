package main

import "golearning/golang/cards"

func main() {

	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")
	cards := deck{"Ace of Diamonds", newCard()}

	ards := newDeck()
	cards.saveToFile("myCards")

	hand, remainignDeck := deal(cards, 5)
	hand.print()
	remainignDeck.print()
	cards.print()
	fmt.Println(cards)

	cards := newDeckFromFile("myCards_")
	fmt.Println(cards)

	cards := newDeck()
	cards = cards.shuffle()
	cards.print()
}

