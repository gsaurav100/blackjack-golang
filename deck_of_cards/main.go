// testing deck package
package main

import (
	"deck_of_cards/deck"
	"fmt"
)

func main() {
	d := deck.New(deck.DefaultSort, deck.AddJokers(2))
	fmt.Println(d, len(d))
}
