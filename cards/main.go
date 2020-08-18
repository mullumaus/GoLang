package main

func main() {
	cards := newDeck() /*define a slice, the size can grow or shrink*/
	//cards.saveToFile("deckfile")
	//newCard := newDeckFromFile("deckfile")
	cards.shuffle()
	cards.print()
}
