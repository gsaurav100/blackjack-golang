package deck

import (
	"math/rand"
	"sort"
	"time"
)

// Suits enum
type Suit int

const (
	Spades Suit = iota
	Diamonds
	Clubs
	Hearts
	JokerSuit
)

// Value enum
type Val int

const (
	_ Val = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	JokerVal
)

var suitsDefault = []Suit{Spades, Diamonds, Clubs, Hearts}
var valDefault = []Val{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}

// printing produces string rather than enum
func (s Suit) String() string {
	switch s {
	case 0:
		return "Spades"
	case 1:
		return "Diamonds"
	case 2:
		return "Clubs"
	case 3:
		return "Hearts"
	case 4:
		return "Joker"
	}
	return "unknown"
}

// printing values
// printing produces string rather than enum
func (v Val) String() string {
	switch v {
	case 1:
		return "Ace"
	case 2:
		return "Two"
	case 3:
		return "Three"
	case 4:
		return "Four"
	case 5:
		return "Five"
	case 6:
		return "Six"
	case 7:
		return "Seven"
	case 8:
		return "Eight"
	case 9:
		return "Nine"
	case 10:
		return "Ten"
	case 11:
		return "Jack"
	case 12:
		return "Queen"
	case 13:
		return "King"
	case 14:
		return "Joker"
	}
	return "unknown"
}

func (c Card) String() string {
	if c.Suit == JokerSuit {
		return "{Joker}"
	}
	return "{" + c.Val.String() + " of " + c.Suit.String() + "}"
}

type Card struct {
	Val  Val
	Suit Suit
}

// function options: https://www.sohamkamani.com/golang/options-pattern/
type DeckOptions func(*[]Card)

// provide a user defined comparison function to execute
func SortDeck(less func(c *[]Card) func(int, int) bool) DeckOptions {
	return func(c *[]Card) {
		sort.Slice(*c, less(c))
	}
}

// example of user defined comparison function - sorts the same as default
func Less(c *[]Card) func(int, int) bool {
	return func(i, j int) bool {
		if (*c)[i].Suit != (*c)[j].Suit {
			return (*c)[i].Suit < (*c)[j].Suit
		}
		return (*c)[i].Val < (*c)[j].Val
	}
}

// of type DeckOptions. Sorts by Numeric val and then suits alphabetically
func DefaultSort(c *[]Card) {
	sort.Slice(*c, func(i, j int) bool {
		if (*c)[i].Suit != (*c)[j].Suit {
			return (*c)[i].Suit < (*c)[j].Suit
		}
		return (*c)[i].Val < (*c)[j].Val
	})
}

// of type DeckOptions. shuffle the deck
func ShuffleDeck(c *[]Card) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(*c), func(i, j int) {
		(*c)[i], (*c)[j] = (*c)[j], (*c)[i]
	})
}

// add Jokers to the deck
func AddJokers(num int) DeckOptions {
	return func(c *[]Card) {
		for i := 0; i < num; i++ {
			joker := Card{Val: JokerVal, Suit: JokerSuit}
			*c = append(*c, joker)
		}
	}
}

// remove certain valued cards. does not maintain order of deck
func RemoveValCards(val Val) DeckOptions {
	return func(c *[]Card) {
		for i := 0; i < len(*c); i++ {
			if (*c)[i].Val == val {
				(*c)[i] = (*c)[len(*c)-1]
				*c = (*c)[:len(*c)-1]
			}
		}
	}
}

// Add num decks
func MultipleDeck(num int) DeckOptions {
	return func(c *[]Card) {
		for i := 0; i < num; i++ {
			newDeck := New()
			*c = append(*c, newDeck...)
		}
	}
}

func New(opts ...DeckOptions) []Card {
	var deck []Card

	// create deck in order from Ace..King and Spades..Hearts
	for _, s := range suitsDefault {
		for _, v := range valDefault {
			deck = append(deck, Card{Val: v, Suit: s})
		}
	}

	// Loop through each option
	for _, opt := range opts {
		opt(&deck)
	}

	return deck
}
