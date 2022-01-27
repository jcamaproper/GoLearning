package main

import "GoLearning/golang/cards"



func main() {

	//cards := []string{"Ace of Diamonds", newCard()}
	//cards = append(cards, "Six of Spades")
	//cards := deck{"Ace of Diamonds", newCard()}

	/* cards := newDeck()
	cards.saveToFile("myCards")

	hand, remainignDeck := deal(cards, 5)
	hand.print()
	remainignDeck.print()
	cards.print()
	fmt.Println(cards) */

	/* cards := newDeckFromFile("myCards_")
	fmt.Println(cards) */

	cards := newDeck()
	cards = cards.shuffle()
	cards.print()
}

/*
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

//create a hand of cards
func deal(d deck, handSize int) (deck, deck) {

	return d[:handSize], d[handSize:]

}

//convert deck []string to string -- join slice of string in one single string
func (d deck) toString() string {

	return strings.Join([]string(d), ",")
}

//write to file
func (d deck) saveToFile(fileName string) error {

	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

//read from a file
func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//slice of byte []byte to string
	d := string(bs)
	//split string into slice of strings []string -- deck
	return strings.Split(d, ",")
}

//swap cards
//Generate random int
func (d deck) shuffle() deck {

	for i := range d {
		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)

		newPosition := r.Intn(len(d) - 1)
		//set a new random position
		d[i], d[newPosition] = d[newPosition], d[i]
	}

	return d
}
*/
