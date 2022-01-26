package main

import "fmt"

// Create a new type of deck
// which is a slice of string
type deck []string

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
			cards = append(cards, suit+" of "+value)
		}
	}
	return cards
}
