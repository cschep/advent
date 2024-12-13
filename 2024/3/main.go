package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
	dCount := 0
	eCount := 0

	// read everything into this and then parse it based on ( or )
	var buf []rune

	// this is changed by the commands do and don't
	mulEnabled := true

	// potential current command
	cmd := ""

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		for _, c := range line {
			if c == '\n' {
				panic("NEW LINE")
			}
			if c == '(' {
				cmd = string(buf)
				fmt.Printf("CMD: [%s] ", cmd)
				if strings.HasSuffix(cmd, "mul") {
					fmt.Print("VALID")
					cmd = "mul"
				} else if strings.HasSuffix(cmd, "do") {
					fmt.Print("VALID")
					cmd = "do"
				} else if strings.HasSuffix(cmd, "don't") {
					fmt.Print("VALID")
					cmd = "don't"
				} else {
					fmt.Printf("invalid!\n")
				}
				buf = []rune{}
			} else if c == ')' {
				if cmd == "mul" {
					inputs := string(buf)
					parts := strings.Split(inputs, ",")
					if len(parts) != 2 {
						panic("WHAT THE FUCK")
					}
					left, right, found := strings.Cut(inputs, ",")
					fmt.Printf(" trying inputs %s", inputs)
					if found {
						leftStr := string(left)
						rightStr := string(right)
						l, lerr := strconv.Atoi(leftStr)
						r, rerr := strconv.Atoi(rightStr)
						if lerr == nil && rerr == nil {
							fmt.Printf(" good inputs!")
							if mulEnabled {
								result += (l * r)
								fmt.Printf(" enabled! adding => %d\n", result)
							} else {
								fmt.Printf(" disabled! not adding.\n")
							}
						} else {
							fmt.Printf(" bad inputs! %s %s %s %s \n", lerr, rerr, leftStr, rightStr)
						}
					}
				} else if cmd == "do" {
					mulEnabled = true
					fmt.Printf(" enabling! inputs: [%s]\n", string(buf))
					eCount++
				} else if cmd == "don't" {
					mulEnabled = false
					fmt.Printf(" disabling! inputs: [%s]\n", string(buf))
					dCount++
				}

				cmd = ""
				buf = []rune{}
			} else {
				buf = append(buf, c)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
	fmt.Println(eCount, dCount)
}
