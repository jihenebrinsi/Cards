package main

func main() {
	//card := newCard()
	//fmt.Println(card)
	//card = "Five of Diamonds"

	//cards := deck{"Ace of Diamonds", newCard()}
	//cards = append(cards, "Six of Spades")
	//cards.print()
	//cards := newDeckFrom("mycards")
	/*	hand, remainingCards := deal(cards, 5)
		hand.print()
		remainingCards.print() */
	//fmt.Println(cards.toString())

	//cards.saveToFile("mycards")
	cards := newDeck()
	cards.shuffle()
	cards.print()

}
