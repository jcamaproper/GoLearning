package main

import "fmt"

// Create a new type of deck
// which is a slice of string
type deck []string

func main() {

	//cards := []string{"Ace of Diamonds", newCard()}
	//cards = append(cards, "Six of Spades")
	//cards := deck{"Ace of Diamonds", newCard()}

	cards := newDeck()
	cards.print()
	fmt.Println(cards)
}

//add new card to the deck
func newCard() string {
	return "Five of Diamonds"
}

// print the content for a deck
func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}
}

//create and return a list of play cards - Slice of string or type deck
func newDeck() deck {

	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}
