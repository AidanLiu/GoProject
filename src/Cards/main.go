package main

func main() {

	cards := newDeck()

	cards.saveToFile("my_cards")
	// hands, lefts := deal(cards, 5)
	// hands.print()
	// fmt.Printf("/n")
	// lefts.print()
	// cards.print()

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	//fmt.Println(cards)
}

func getCard() string {
	return "5 of diamons"
}
