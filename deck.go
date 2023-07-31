package main

//new type deck
//which is slice of string
import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := deck{"Heart", "Spade", "Diamond", "Club"}
	cardValues := deck{"satu", "dua", "tiga", "empat"}

	for _, Suits := range cardSuits {
		for _, Values := range cardValues {
			cards = append(cards, Values+" dari "+Suits)
		}
	}
	return cards
}

func deal(mydeck deck, handsize int) (deck, deck) {
	return mydeck[handsize:], mydeck[:handsize]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) print() {
	for i, cards := range d {
		fmt.Println(i, cards)
	}
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckfromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
