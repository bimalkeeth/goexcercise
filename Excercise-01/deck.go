package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type deck []string

func newDeck() deck{
	cards:=deck{}
	cardSuits:=[]string{"Spades","Diamonds","Hearts","Clubs"}
    cardValues:=[]string{"Ace","Two","Three","Four"}

    for _,suit :=range cardSuits{
    	for _,value:=range cardValues{

    		cards=append(cards,value + " of " + suit )
		}
	}
	return cards
}

func deal(d deck,handSize int)(deck,deck){
   return d[:handSize],d[handSize:]
}

func(dec deck) print(){
	for i,dec :=range dec {
		fmt.Println(i,dec)
	}
}

func (dec deck)toString()string{
	 str:= []string(dec)
	 strJoin:= strings.Join(str,",")
	 return strJoin
}

func(dec deck)saveToFile(fileName string)error{
 return ioutil.WriteFile(fileName,[]byte(dec.toString()),0666)

}
func newDeckFromFile(fileName string) (deck,error){
	bs,err:=ioutil.ReadFile(fileName)
	if err!=nil || len(bs)>0{
	   fmt.Println(err)
	   return nil,err
	}
	cards:=strings.Split(string(bs),",")
	return deck(cards),nil
}

func(dec deck)shuffle(){
	rand.Seed(time.Now().UTC().Unix())
	for i:=range dec{
       newPosition:=rand.Intn(len(dec)-1)
       dec[i],dec[newPosition]=dec[newPosition],dec[i]
	}
}