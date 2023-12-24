package main

import (
	"aoc/util"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("inputs/14.input.small")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	grids := []util.Grid{}
	g := util.Grid{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			grids = append(grids, g)
			g = util.Grid{}
		} else {
			g = append(g, line)
		}
	}
	grids = append(grids, g)

	g.Print()
	tiltNorth(g)
	g.Print()

	fmt.Println(load(g))
}

func tiltNorth(g util.Grid) {
	for x := 0; x < g.Width(); x++ {
		nextMove := -1
		for y := 0; y < g.Height(); y++ {
			tile := g.Get(x, y)
			if tile == "." && nextMove == -1 {
				nextMove = y
			}
			if tile == "#" {
				nextMove = y + 1
			}
			if tile == "O" {
				if nextMove > -1 {
					//move it
					g.Set(x, y, ".")
					g.Set(x, nextMove, "O")
					//the next move is the spot we just created
					nextMove = nextMove + 1
				}
			}
		}
	}
}

func load(g util.Grid) int {
	sum := 0
	g.Each(func(s string, x, y int) {
		if s == "O" {
			sum += (g.Height() - y)
		}
	})
	return sum
}
