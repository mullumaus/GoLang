package main

import "fmt"

func defSlice() {
	//intvalue := []int{}
	cards := []string{"Ace of diamonds", newCard()} /*define a slice, the size can grow or shrink*/
	cards = append(cards, "Six of spades")
	/* append does not modify exisitng slice, it creates a new slice and assign to cards */

	fmt.Println(cards)
	for i, card := range cards {
		fmt.Println(i, card)
	}
}
