package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	file, err := os.Open("4.input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	countMap := map[int]int{}

	result := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		result++
		cardNumStr, gameStr, found := strings.Cut(line, ":")
		if !found {
			panic("NOPE")
		}
		var cardNum int
		_, err := fmt.Sscanf(cardNumStr, "Card %d", &cardNum)
		if err != nil {
			panic(err)
		}

		mineStr, winningStr, found := strings.Cut(gameStr, "|")
		mineStr = strings.Trim(mineStr, " ")
		winningStr = strings.Trim(winningStr, " ")
		if !found {
			panic("not found")
		}

		mine, winning := strings.Split(mineStr, " "), strings.Split(winningStr, " ")
		cardScore := 0
		fmt.Println(cardNum, mine, winning)
		for _, v := range mine {
			if v == "" {
				continue
			}
			if err != nil {
				panic(err)
			}

			win := contains(winning, v) //🤘🏻
			if win {
				fmt.Println(v, win)
				cardScore++
			}
		}

		multiplier := countMap[cardNum] + 1
		fmt.Println("multi ->", multiplier)

		for i := 1; i < cardScore+1; i++ {
			countMap[cardNum+i] += multiplier
		}

		fmt.Println(countMap)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, v := range countMap {
		result += v
	}

	fmt.Println(result)
}
