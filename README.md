# Blackjack game

## Introduction

This is a program created in Golang to play blackjack against a bot. The game is played in the CLI and requires Golang to be installed. [Install Golang](#install-golang)

The aim of the game is to get the sum of the cards as close to 21 as possible. For further information read [Rules of the Game](#rules-of-the-game)

To run the program, execute the following command in the root folder:
> go run main.go

## Rules of the Game

The rules of blackjack are very simple:

1. You and the dealer are both dealt 2 cards each
2. You are able to view 1 of the dealers cards while the dealer cannot see either of your cards
3. The aim of the game is to make the sum of the values of your cards as close to 21 as possible. The suits have no meaning in this game
4. The value of each card is as follows:
    - All numbered cards (2-10) have their value 
    - All face cards (Jack, Queen, and King) are worth 10
    - An Ace is worth either 1 or 11 - the player can choose which one gets them closer to 21
5. The user has 2 options on their turn
    - HIT (H): They get given another card from the deck
    - STAY (S): They end their turn and pass over to the dealer. You can no longer get new cards after this
6. If a player goes over 21, they go BUST and lose the game
7. If neither player goes BUST, then the player closest to 21 wins the game

## Install Golang

Use the following guide to install golang. Choose your appropriate Opertating System

https://go.dev/doc/install

Ensure you have golang installed by running
> go version

