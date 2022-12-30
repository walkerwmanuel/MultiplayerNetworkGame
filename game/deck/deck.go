package deck

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func (d *Deck) Shuffle() {
	if len(d.Cards) == 0 {
		return
	}

	// iterate through each card in Cards of deck
	for i := range d.Cards {
		// pull out card of current iteration index 0...len(d.Cards)
		card := d.Cards[i]
		// creates a new (random) position from 0...len(d.Cards)
		newPos, _ := rand.Int(rand.Reader, big.NewInt(int64(len(d.Cards))))
		// convert newPos to int
		newPosInt := newPos.Uint64()
		// pull out card in new position
		otherCard := d.Cards[newPosInt]
		// swap them
		d.Cards[i] = otherCard
		d.Cards[newPosInt] = card
	}
}

// Deck.Print - prints content of a deck
func (d *Deck) Print() {
	fmt.Println()
	for i := range d.Cards {
		fmt.Printf("[%s:%s] ", d.Cards[i].Suite, d.Cards[i].Value)
	}
	fmt.Println()
}

// NewCard - makes a new card
func NewCard(v, s string) Card {
	c := Card{
		Suite: s,
		Value: v,
	}
	return c
}

// NewDeck - creates a deck with 52 cards
func NewDeck() *Deck {
	d := Deck{}
	d.Cards = make([]Card, 52)
	index := 0
	for i := range suites {
		for j := range values {
			d.Cards[index] = NewCard(values[j], suites[i])
			index++
		}
	}
	return &d
}
