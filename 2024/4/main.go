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
		// fmt.Println("scanning at", x, y)
		for j := -1; j <= 1; j++ {
			for i := -1; i <= 1; i++ {
				// the idea here is to pick a direction
				// then loop through "walking" in that direction
				for k := 1; k < 4; k++ {
					newX := x + (i * k)
					newY := y + (j * k)

					fmt.Println("checking", x, y, i, j, newX, newY)
					if g.Get(newX, newY) == target {
						if target == "M" {
							fmt.Println("found M at", newX, newY)
							target = "A"
						} else if target == "A" {
							fmt.Println("found A at", newX, newY)
							target = "S"
						} else if target == "S" {
							fmt.Println("found S at", newX, newY)
							result++
							return
						}
					} else {
						fmt.Println("breaking")
						break
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
	//
	// fmt.Println(g.Get(0, 0))
	// fmt.Println(g.Get(0, 1))
	// fmt.Println(g.Get(0, 2))
	// fmt.Println(g.Get(0, 3))

	// scan(4, 1, "M")

	fmt.Println(result)
}
