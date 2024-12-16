package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("2024/1/1.input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lefts := []int{}
	rights := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		left, right, found := strings.Cut(line, "   ")
		if !found {
			panic("input bad")
		}

		l, err := strconv.Atoi(left)
		if err != nil {
			panic(err)
		}

		r, err := strconv.Atoi(right)
		if err != nil {
			panic(err)
		}

		lefts = append(lefts, l)
		rights = append(rights, r)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	part1 := part1(lefts, rights)
	log.Println(part1)

	part2 := part2(lefts, rights)
	log.Println(part2)
}

func part2(lefts, rights []int) int {
	result := 0

	counts := map[int]int{}
	for _, v := range rights {
		counts[v] += 1
	}

	for _, v := range lefts {
		count := counts[v]
		result += v * count
	}

	return result
}

func part1(lefts, rights []int) int {
	sort.Slice(lefts, func(i, j int) bool {
		return lefts[i] < lefts[j]
	})

	sort.Slice(rights, func(i, j int) bool {
		return rights[i] < rights[j]
	})

	result := 0
	for i := 0; i < len(lefts); i++ {
		diff := abs(lefts[i] - rights[i])
		result += diff
	}

	return result
}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}
