package main

import "fmt"

func defvar() {
	var card string = "Ace of Spades"
	fmt.Println(card)
	/* card := newCard()
	var card string = "Ace of Spades" and card := "Ace of Spades"  these two statements are the same
	only use ":=" in var definition/initialization, for new variable only */

}
func printState() {
	fmt.Println("California")
}
