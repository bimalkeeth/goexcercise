package main

func main() {
	 cards :=newDeck()
	 //cards.print()
	 //hand,remainingCards:= deal(cards,5)
	 //hand.print()
	 //remainingCards.print()
	 //fmt.Println(hand.toString())
	 //fmt.Println(remainingCards.toString())
     //err:=cards.saveToFile("my_cards")
     //if err!=nil{
     	//fmt.Println("err")
	 //}
	 cards.shuffle()
	 cards.print()

}
