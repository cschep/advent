package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func numbers(s string) []int {
	var n []int
	for _, f := range strings.Fields(s) {
		i, err := strconv.Atoi(f)
		if err == nil {
			n = append(n, i)
		}
	}
	return n
}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}

func process(levels []int) bool {
	if len(levels) < 2 {
		return false
	}

	dir := 1
	if levels[1] < levels[0] {
		dir = -1
	}

	prev := levels[0]
	for _, v := range levels[1:] {
		if prev < v && dir == -1 {
			return false
		}
		if prev > v && dir == 1 {
			return false
		}

		diff := abs(prev - v)
		if diff > 3 || diff < 1 {
			return false
		}

		prev = v
	}

	return true
}

func part1(input [][]int) int {
	result := 0
	for _, levels := range input {
		if process(levels) {
			result += 1
		}
	}

	return result
}

// B R U T E F O R C E
func part2(input [][]int) int {
	result := 0

	for _, levels := range input {
		// if it is safe, move on
		if process(levels) {
			result += 1
			continue
		}

		// if it is not safe, now try brute forcing each level out until one is safe
		for i := 0; i < len(levels); i++ {
			newLevels := make([]int, 0, len(levels)-1)
			newLevels = append(newLevels, levels[:i]...)
			newLevels = append(newLevels, levels[i+1:]...)

			if process(newLevels) {
				result += 1
				break
			}
		}
	}

	return result
}

func main() {
	file, err := os.Open("inputs/2.input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		levels := numbers(line)
		input = append(input, levels)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// fmt.Println(part1(input))
	fmt.Println(part2(input))
}
