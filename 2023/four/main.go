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

	result := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, ":")
		mineStr, winningStr, found := strings.Cut(numbers[1], "|")
		mineStr = strings.Trim(mineStr, " ")
		winningStr = strings.Trim(winningStr, " ")
		if !found {
			panic("not found")
		}
		mine, winning := strings.Split(mineStr, " "), strings.Split(winningStr, " ")
		cardScore := 0
		fmt.Println(mine, winning)
		for _, v := range mine {
			if v == "" {
				continue
			}
			if err != nil {
				panic(err)
			}

			win := contains(winning, v)
			if win {
				fmt.Println(v, win)
				if cardScore == 0 {
					cardScore = 1
				} else {
					cardScore = 2 * cardScore
				}
			}
		}
		fmt.Println("adding", cardScore)
		result += cardScore
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
