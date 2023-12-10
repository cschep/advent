package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var maxes = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func isPossible(line string) (int, bool) {
	l := strings.Split(line, ":")
	gameID, err := strconv.Atoi(strings.Split(l[0], " ")[1])
	if err != nil {
		panic(err)
	}
	sets := strings.Split(l[1], ";")
	fmt.Println("gameID:", gameID)
	possible := true
	for _, set := range sets {
		for _, v := range strings.Split(set, ",") {
			var color string
			var amount int
			fmt.Sscanf(v, "%d %s", &amount, &color)
			if maxes[color] < amount {
				fmt.Println("IMPOSSIBLE", maxes[color])
				possible = false
			}
		}
	}

	return gameID, possible
}

func minPower(line string) int {
	l := strings.Split(line, ":")
	soFar := map[string]int{"red": 0, "blue": 0, "green": 0}
	sets := strings.Split(l[1], ";")

	for _, set := range sets {
		for _, v := range strings.Split(set, ",") {
			var color string
			var amount int
			fmt.Sscanf(v, "%d %s", &amount, &color)
			if soFar[color] < amount {
				soFar[color] = amount
			}
		}
	}

	return soFar["red"] * soFar["green"] * soFar["blue"]
}

func main() {
	file, err := os.Open("inputs/2.input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// <ID>: [GAME ENTRIES]; [GAME ENTRIES]
	// split on colon to get ID
	// split on ; to get rounds?
	// split on , to get each throw
	// split on space to get color -> amount
	result := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		//2a
		// gameID, possible := isPossible(line)

		// if possible {
		// 	result += gameID
		// }

		//2b
		power := minPower(line)
		result += power
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(result)
}
