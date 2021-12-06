package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func part1() {
	file, err := os.Open("2.in")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var total int = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		vals := strings.Split(scanner.Text(), "x")
		ints := []int{}

		for _, val := range vals {
			new, err := strconv.Atoi(val)
			if err != nil {
				panic("DEAD")
			}
			ints = append(ints, new)
		}

		l, w, h := ints[0], ints[1], ints[2]
		side1 := l * w
		side2 := l * h
		side3 := w * h

		sides := []int{side1, side2, side3}
		smallestSide := sides[0]
		for _, side := range sides {
			if side < smallestSide {
				smallestSide = side
			}
		}

		total += (2 * side1) + (2 * side2) + (2 * side3) + smallestSide
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println(total)
}

func part2() {
	file, err := os.Open("2.in")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var total int = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		vals := strings.Split(scanner.Text(), "x")
		ints := []int{}

		for _, val := range vals {
			new, err := strconv.Atoi(val)
			if err != nil {
				panic("DEAD")
			}
			ints = append(ints, new)
		}

		sort.Ints(ints)
		perim := (2 * ints[0]) + (2 * ints[1])
		vol := ints[0] * ints[1] * ints[2]
		total += perim + vol
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println(total)
}

func main() {
	part1()
	part2()
}
