package main

import (
	"webblackjack/game/deck"
)

func main() {
	d := deck.NewDeck()
	d.Print()
	d.Shuffle()
	d.Print()
}
