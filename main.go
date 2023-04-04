package main

import (
	"blackjack/deck_of_cards/deck"
	"fmt"
	"sort"
	"strings"
)

const WINNING_TARGET = 21

// recursively get all possible scores (multiple scores with Ace)
func doScoreRecursion(hand []deck.Card, currScore int, allScores *[]int) {
	if len(hand) == 0 {
		*allScores = append(*allScores, currScore)
		return
	}
	c := hand[len(hand)-1]
	if c.Val == deck.King || c.Val == deck.Queen || c.Val == deck.Jack {
		doScoreRecursion(hand[:len(hand)-1], currScore+10, allScores)
	} else if c.Val == deck.Ace {
		doScoreRecursion(hand[:len(hand)-1], currScore+1, allScores)
		doScoreRecursion(hand[:len(hand)-1], currScore+11, allScores)
	} else {
		doScoreRecursion(hand[:len(hand)-1], currScore+int(c.Val), allScores)
	}
}

// get the score for the provided hand
func getScore(hand []deck.Card) []int {
	var score []int
	doScoreRecursion(hand, 0, &score)
	return score
}

// check if all possible user scores are above WINNING_TARGET
func isUserBust(hand []deck.Card) bool {
	scores := getScore(hand)
	for _, s := range scores {
		if s <= WINNING_TARGET {
			return false
		}
	}
	return true
}

// only print out valid scores (i.e. scores below WINNING_TARGET)
func printValidScores(hand []deck.Card) {
	scores := getScore(hand)
	fmt.Print("[")
	for _, s := range scores {
		if s > WINNING_TARGET {
			continue
		}
		fmt.Print(" ")
		fmt.Print(s)
		fmt.Print(" ")
	}
	fmt.Print("]")
	fmt.Println()
}

// get highest possible score - returns -1 if all scores are bust
func getHighestValidScore(hand []deck.Card) int {
	scores := getScore(hand)
	sort.Ints(scores)

	// all scores are bust
	if scores[0] > WINNING_TARGET {
		return -1
	}

	// return highest score before bust score
	for i, s := range scores {
		if s > WINNING_TARGET {
			return scores[i-1]
		}
	}

	// if no scores are bust then last element is highest
	return scores[len(scores)-1]
}

// deal a card from provided deck to provided player hand
func deal(d *[]deck.Card, hand *[]deck.Card) {
	*hand = append(*hand, (*d)[len(*d)-1]) // give card from top of deck to hand
	*d = (*d)[:len(*d)-1]                  // remove card from deck
}

// run dealer logic on this turn
func playDealer(d *[]deck.Card, hand *[]deck.Card) {
	for {
		score := getHighestValidScore(*hand)
		// dealer is bust
		if score == -1 {
			break
		}
		hasAce := false
		for _, c := range *hand {
			if c.Val == deck.Ace {
				hasAce = true
			}
		}

		if hasAce {
			if score <= 17 {
				fmt.Println(*hand)
				deal(d, hand)
			} else {
				break
			}
		} else {
			if score <= 16 {
				deal(d, hand)
			} else {
				break
			}
		}
	}

}

// print deck
func printHand(d []deck.Card) {
	for _, c := range d {
		fmt.Print(c, " ")
	}
	fmt.Println()
}

func main() {

	d := deck.New(deck.ShuffleDeck)

	var playerCards []deck.Card
	var dealerCards []deck.Card

	// deal 2 cards to the player and the dealer
	deal(&d, &playerCards)
	deal(&d, &dealerCards)
	deal(&d, &playerCards)
	deal(&d, &dealerCards)

	// show user the one of the dealers cards
	fmt.Println("DEALERS CARDS")
	fmt.Println(dealerCards[0], "[HIDDEN]")
	fmt.Println("-------------------")
	fmt.Println("YOUR CARDS")
	printHand(playerCards)
	fmt.Print("Your score(s): ")
	printValidScores(playerCards)

	var userInput string
	for {
		// Ask user to hit or stay
		fmt.Println()
		fmt.Print("HIT (H) or STAY (S): ")
		fmt.Scanln(&userInput)
		fmt.Println()
		if strings.ToUpper(userInput) == "H" {
			// deal the user another card
			deal(&d, &playerCards)

			// if user is busted end the game
			if isUserBust(playerCards) {
				fmt.Println("*******YOU WENT BUST*******")
				fmt.Println("YOUR CARDS")
				printHand(playerCards)
				fmt.Println("Your score(s): ", getScore(playerCards)) // show all scores
				fmt.Println()
				fmt.Println("*******************************")
				fmt.Println("***********YOU LOST************")
				fmt.Println("*******************************")
				break
			}

			fmt.Println("DEALERS CARDS")
			fmt.Println(dealerCards[0], "[HIDDEN]")
			fmt.Println("-------------------")
			fmt.Println("YOUR CARDS")
			printHand(playerCards)
			fmt.Print("Your score(s): ")
			printValidScores(playerCards) // only show scores less than WINNING_TARGET
		} else if strings.ToUpper(userInput) == "S" {
			// dealer makes decisions
			playDealer(&d, &dealerCards)

			// check if dealer bust - player wins instantly
			if isUserBust(dealerCards) {
				fmt.Println("*******DEALER WENT BUST*******")
				fmt.Println("DEALERS CARDS")
				printHand(dealerCards)
				fmt.Println("Dealer score(s): ", getScore(dealerCards))
				fmt.Println("-------------------")
				fmt.Println("YOUR CARDS")
				printHand(playerCards)
				fmt.Println("Your score: ", getHighestValidScore(playerCards)) // show all scores
				fmt.Println()
				fmt.Println("*******************************")
				fmt.Println("************YOU WON************")
				fmt.Println("*******************************")
				break
			}

			// show all the dealers cards
			dealerScore := getHighestValidScore(dealerCards)
			playerScore := getHighestValidScore(playerCards)
			fmt.Println("DEALERS CARDS")
			printHand(dealerCards)
			fmt.Println("Dealer score: ", dealerScore)
			fmt.Println("-------------------")
			fmt.Println("YOUR CARDS")
			printHand(playerCards)
			fmt.Println("Your score: ", playerScore)
			if playerScore > dealerScore {
				fmt.Println()
				fmt.Println("*******************************")
				fmt.Println("************YOU WON************")
				fmt.Println("*******************************")
			} else {
				fmt.Println()
				fmt.Println("*******************************")
				fmt.Println("***********YOU LOST************")
				fmt.Println("*******************************")
			}

			// end the game
			break
		} else {
			fmt.Println("Invalid input.")
		}
	}
}
