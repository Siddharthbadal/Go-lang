package main

import "fmt"

func main() {
	cards := []string{newCard(), "Ace of Queens"}
	cards = append(cards, "Six of Spades")
	for i, card := range cards {
		fmt.Println(i, card)
	}
	//fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
