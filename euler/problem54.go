package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Problem54 struct{}

type PokerHand struct {
	Cards [5]Card
}

type Card struct {
	CardNumber int
	CardType   string
}

var (
	cardNumbers = map[string]int{
		"2": 0, "3": 1, "4": 2, "5": 3, "6": 4, "7": 5, "8": 6, "9": 7,
		"T": 8, "J": 9, "Q": 10, "K": 11, "A": 12}
)

func (p *Problem54) GetSolution() string {

	pokerHandPairs := readPokerHands("problem54.txt")
	wins := 0
	for _, pair := range pokerHandPairs {
		if getWinner(pair[0], pair[1]) == 1 {
			wins++
		}
	}
	return fmt.Sprint(wins)
}

func getWinner(h1 PokerHand, h2 PokerHand) int {
	r1, v1 := getRankValue(h1)
	r2, v2 := getRankValue(h2)
	if r1 > r2 {
		return 1
	}
	if r1 < r2 {
		return 2
	}
	if v1 > v2 {
		return 1
	}
	return 2
}

func getRankValue(h PokerHand) (int, int) {
	if isRoyalFlush(h) {
		return 10, 14
	}
	ok, max := isStraightFlush(h)
	if ok {
		return 9, max
	}

	ok, max = isFourOfKind(h)
	if ok {
		return 8, max
	}
	fh, max1, max2 := isFullHouse(h)
	if fh {
		return 7, max1 + max2
	}

	if isFlush(h) {
		return 6, getHighCard(h)
	}
	ok, max = isStraight(h)
	if ok {
		return 5, max
	}
	ok, max = isThreeOfAKind(h)
	if ok {
		return 4, max
	}
	ok, max1, max2 = isTwoPairs(h)
	if ok {
		return 3, max1 + max2
	}
	ok, max = isOnePair(h)
	if ok {
		return 2, max
	}
	return 1, getHighCard(h)
}

func getHighCard(h PokerHand) int {
	cardValues := make([]int, 14)
	for _, card := range h.Cards {
		cardValues[card.CardNumber]++
	}
	for i := 13; i > 0; i-- {
		if cardValues[i] > 0 {
			return i
		}
	}

	return 0
}

func isOnePair(h PokerHand) (bool, int) {
	cardValues := make([]int, 14)
	for _, card := range h.Cards {
		cardValues[card.CardNumber]++
	}

	for i := 13; i > 0; i-- {
		if cardValues[i] > 1 {
			return true, i
		}
	}

	return false, -1
}

func isThreeOfAKind(h PokerHand) (bool, int) {
	cardValues := make([]int, 14)
	for _, card := range h.Cards {
		cardValues[card.CardNumber]++
	}

	for i := 13; i > 0; i-- {
		if cardValues[i] > 2 {
			return true, i
		}
	}

	return false, -1
}

func isTwoPairs(h PokerHand) (bool, int, int) {
	cardValues := make([]int, 14)
	for _, card := range h.Cards {
		cardValues[card.CardNumber]++
	}
	pair1 := -1
	for i := 13; i > 0; i-- {
		if cardValues[i] > 1 {
			if pair1 < 0 {
				pair1 = i
				continue
			}
			return true, pair1, i
		}
	}

	return false, -1, -1
}

func isStraight(h PokerHand) (bool, int) {
	cardValues := make([]bool, 14)
	for _, card := range h.Cards {
		cardValues[card.CardNumber] = true
	}

	consecCount := 0
	for i, v := range cardValues {
		if !v && consecCount > 0 {
			return consecCount == 5, i
		}
		if v {
			consecCount++
		}
	}
	return consecCount == 5, 13
}

func isFullHouse(h PokerHand) (bool, int, int) {
	cardValues := make([]int, 14)
	for _, card := range h.Cards {
		cardValues[card.CardNumber]++
	}
	pair1 := -1
	pair2 := -1
	for i := 13; i > 0; i-- {
		if cardValues[i] == 3 {
			pair1 = i
		} else if cardValues[i] == 2 {
			pair2 = i
		}
	}

	return pair1 > 0 && pair2 > 0, pair1, pair2
}

func isFourOfKind(h PokerHand) (bool, int) {
	cardValues := make([]int, 14)
	for _, card := range h.Cards {
		cardValues[card.CardNumber]++
	}
	for i := 13; i > 0; i-- {
		if cardValues[i] == 4 {
			return true, i
		}
	}

	return false, -1
}

func isFlush(h PokerHand) bool {
	suit := make(map[string]bool)

	for _, card := range h.Cards {
		suit[card.CardType] = true
	}

	return len(suit) == 1
}

func isStraightFlush(h PokerHand) (bool, int) {
	if !isFlush(h) {
		return false, -1
	}

	return isStraight(h)
}

func isRoyalFlush(h PokerHand) bool {
	isStraight, startingIndex := isStraightFlush(h)
	return isStraight && startingIndex == 9
}

func readPokerHands(filePath string) [][]PokerHand {
	content, _ := ioutil.ReadFile(filePath)

	lines := strings.Split(string(content), "\n")
	pokerHands := make([][]PokerHand, 0)
	for _, line := range lines {
		cards := strings.Split(line, " ")
		ph1 := buildPokerHand(cards[0:5])
		ph2 := buildPokerHand(cards[5:])
		pokerHands = append(pokerHands, []PokerHand{ph1, ph2})
	}

	return pokerHands
}

func buildPokerHand(cards []string) PokerHand {
	ph := PokerHand{}
	for i, c := range cards {
		symbols := strings.Split(c, "")
		ph.Cards[i] = Card{
			CardNumber: cardNumbers[symbols[0]],
			CardType:   symbols[1]}
	}
	return ph
}
