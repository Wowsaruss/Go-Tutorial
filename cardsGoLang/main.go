package main

func main() {
	cards := newDeck()
	cards.saveToFile("new_deck")
	cards.shuffle()
	cards.saveToFile("shuffled_deck")
	hand, remainingDeck := deal(cards, 7)

	remainingDeck.print()
	hand.print()
}
