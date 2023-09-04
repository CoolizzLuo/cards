package main

func main() {
	cards := newDeck()
	cards.shuffle()

	cards.print()

	deal(cards, 5)

}
