package main

import (
	"aoc/util"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("2024/4/4.input.small")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	g := util.NewGrid(0, 0)
	for scanner.Scan() {
		line := scanner.Text()
		g.PushRow(line)
	}

	g.Print()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	result := 0
	var scan func(x, y int, target string)
	scan = func(x, y int, target string) {
		for j := y - 1; j <= y+1; j++ {
			for i := x - 1; i <= x+1; i++ {
				if g.Get(i, j) == target {
					if target == "M" {
						scan(i, j, "A")
					} else if target == "A" {
						scan(i, j, "S")
					} else if target == "S" {
						result++
					}
				}
			}
		}
	}

	g.Each(func(s string, x int, y int) {
		// for each letter that is an X, start a walk!
		if s == "X" {
			scan(x, y, "M")
		}
	})

	fmt.Println(result)
}
