package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func part1() {
	f, err := os.Open("1.in")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	count := 0
	scanner.Scan()
	prev, _ := strconv.Atoi(scanner.Text())
	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())
		if val > prev {
			count++
		}
		prev = val
	}

	fmt.Println(count)
}

func part2() {
	f, err := os.Open("1.in")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	count := 0
	scanner.Scan()
	ppprev, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	pprev, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	prev, _ := strconv.Atoi(scanner.Text())

	lastSum := ppprev + pprev + prev
	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())
		sum := pprev + prev + val
		if sum > lastSum {
			count++
		}
		pprev = prev
		prev = val
		lastSum = sum
	}

	fmt.Println(count)
}

func main() {
	part1()
	part2()
}
