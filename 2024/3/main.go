package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func part1(input string) {
	fmt.Println(input)
}

func main() {
	file, err := os.Open("inputs/3.input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := 0
	for scanner.Scan() {
		line := scanner.Text()
		// prev := rune(line[0])

		//anything but an advanced state i think
		state := 'z'

		var left, right []rune
		for _, c := range line {
			if c == 'm' {
				state = 'm'
			} else if c == 'u' && state == 'm' {
				state = 'u'
			} else if c == 'l' && state == 'u' {
				state = 'l'
			} else if c == '(' && state == 'l' {
				state = '('
			} else if unicode.IsDigit(c) && (state == '(' || state == 'x') {
				left = append(left, c)
				state = 'x'
			} else if c == ',' && state == 'x' {
				state = ','
			} else if unicode.IsDigit(c) && (state == ',' || state == 'y') {
				right = append(right, c)
				state = 'y'
			} else if c == ')' && state == 'y' {
				leftStr := string(left)
				rightStr := string(right)

				fmt.Println(leftStr, rightStr)

				l, err := strconv.Atoi(leftStr)
				if err != nil {
					panic(err)
				}
				r, err := strconv.Atoi(rightStr)
				if err != nil {
					panic(err)
				}
				res := l * r
				fmt.Printf("found %d * %d = %d\n", l, r, res)
				result += res

				state = 'z'
				left = []rune{}
				right = []rune{}
			} else {
				//reset
				state = 'z'
				left = []rune{}
				right = []rune{}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
