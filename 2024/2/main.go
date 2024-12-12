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

func main() {
	file, err := os.Open("inputs/2.input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		levels := numbers(line)
		if process(levels) {
			result += 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
