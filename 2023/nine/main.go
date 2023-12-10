package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("9.input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	result := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		numberStrings := strings.Split(line, " ")

		lineNumbers := []int{}
		for _, v := range numberStrings {
			i, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}

			lineNumbers = append(lineNumbers, i)
		}

		ds := [][]int{}
		d := diffs(lineNumbers)
		for {
			ds = append(ds, d)
			if all(d, 0) {
				break
			}
			d = diffs(d)
		}

		//part 1
		nextStep := 0
		// for i := len(ds) - 1; i >= 0; i-- {
		// 	s := ds[i][len(ds[i])-1]
		// 	nextStep += s
		// }

		//part 2
		for i := len(ds) - 1; i >= 0; i-- {
			s := ds[i][0]
			nextStep = s - nextStep
		}

		answer := lineNumbers[0] - nextStep
		result += answer
	}

	fmt.Println()
	fmt.Println(result)
}

// GENERICS!!
func all[T comparable](a []T, same T) bool {
	for _, v := range a {
		if v != same {
			return false
		}
	}

	return true
}

func diffs(a []int) []int {
	prev := a[0]
	diffs := []int{}
	for i := 1; i < len(a); i++ {
		curr := a[i]
		diff := curr - prev
		diffs = append(diffs, diff)
		prev = curr
	}

	return diffs
}
