package main

import (
	"aoc/util"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("inputs/14.input")
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

	// g.Print()
	// tiltNorth(g)
	// g.Print()

	cycles := 1000
	loads := []int{}
	for i := 0; i < cycles; i++ {
		for i := 0; i < 4; i++ {
			tiltNorth(g)
			g = g.RotateCW()
		}
		loads = append(loads, load(g))
	}

	found := false
	var from, to int
	// cycleLength := 0
	for i := 100; i < len(loads); i++ {
		//find the next instance of it
		if found {
			break
		}
		for j := i + 1; j < len(loads); j++ {
			if loads[i] == loads[j] {
				if checkCycle(i, j, loads) {
					from, to = i, j
					found = true
					break
				}
			}
		}
	}

	if found {
		fmt.Println("cycle found", from, to)
		distance := to - from
		for i := from; i < to; i++ {
			fmt.Println(i, loads[i])
		}

		megaCycles := 1000000000
		loadsIndex := (megaCycles - from - 1) % distance
		fmt.Println(loadsIndex, loads[from+loadsIndex])
	}
}

func checkCycle(from, to int, loads []int) bool {
	for i := 0; i < len(loads); i++ {
		if to+i < len(loads) && loads[from+i] != loads[to+i] {
			return false
		}
	}
	return true
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
