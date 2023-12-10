package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type Hand struct {
	cards string
	bid   int
}

func main() {
	file, err := os.Open("inputs/7.input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var hands []Hand

	// round 1
	// "AKQJT98765432"
	// round 2
	// "AKQT98765432J"
	cardValues := map[byte]int{
		'A': 13, 'K': 12, 'Q': 11, 'T': 10, '9': 9, '8': 8, '7': 7, '6': 6, '5': 5, '4': 4, '3': 3, '2': 2, '1': 1, 'J': 0,
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var cards string
		var bid int
		_, err := fmt.Sscanf(line, "%s %d", &cards, &bid)
		if err != nil {
			panic(err)
		}

		h := Hand{cards: cards, bid: bid}
		hands = append(hands, h)
	}

	sort.Slice(hands, func(i, j int) bool {
		leftHand, rightHand := hands[i], hands[j]

		leftType, _ := bestType(leftHand.cards)
		rightType, _ := bestType(rightHand.cards)

		if leftType == rightType {
			for k := 0; k < len(leftHand.cards); k++ {
				l, r := cardValues[leftHand.cards[k]], cardValues[rightHand.cards[k]]
				if l != r {
					return l < r
				}

				if l == r && k == len(leftHand.cards)-1 {
					panic("THIS IS NOT GOOD")
				}
			}
		}

		return leftType < rightType
	})

	result := 0
	for i, hand := range hands {
		rank := i + 1
		result += hand.bid * rank
	}

	for _, v := range hands {
		fmt.Println(v)
	}

	fmt.Println(result)
}

func bestType(cards string) (int, string) {
	// if a hand has a J, rotate the J's through all values to find the best type?
	if !strings.Contains(cards, "J") {
		return getType(cards), ""
	}

	best := 0
	bestSub := ""
	// loop through the values replacing J's
	for _, v := range "AKQT98765432" {
		str := string(v)
		newCards := strings.ReplaceAll(cards, "J", str)
		newType := getType(newCards)

		if newType > best {
			best = newType
			bestSub = str
		}
	}

	// fmt.Println(best, bestSub, ":", cards, " -> ", strings.ReplaceAll(cards, "J", bestSub))
	return best, bestSub
}

func getType(cards string) int {
	cardMap := map[rune]int{}
	for _, card := range cards {
		cardMap[card] += 1
	}

	hasFive := false
	hasFour := false
	hasThree := false
	hasTwo := false
	hasTwoTwice := false
	for _, v := range cardMap {
		if v == 5 {
			hasFive = true
			break
		}

		if v == 4 {
			hasFour = true
			break
		}

		if v == 3 {
			hasThree = true
		}

		if v == 2 && hasTwo {
			hasTwoTwice = true
		}

		if v == 2 && !hasTwo {
			hasTwo = true
		}
	}

	if hasFive {
		return 6
	} else if hasFour {
		return 5
	} else if hasThree && hasTwo {
		return 4
	} else if hasThree && !hasTwo {
		return 3
	} else if hasTwo && hasTwoTwice {
		return 2
	} else if hasTwo && !hasTwoTwice {
		return 1
	}

	return 0
}
