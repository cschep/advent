package main

import (
	. "aoc/util"
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("inputs/11.input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	g := NewGrid(0, 0)
	for scanner.Scan() {
		line := scanner.Text()
		g.PushRow(line)
	}

	g.Print()
	expand(g)
	g.Print()

	gals := findGalaxies(*g)
	result := map[string]int{}
	for i1, g1 := range gals {
		for i2, g2 := range gals {
			if g1 == g2 {
				continue
			}
			key := fmt.Sprintf("%d<->%d", i1+1, i2+1)
			key2 := fmt.Sprintf("%d<->%d", i2+1, i1+1)
			if result[key] != 0 || result[key2] != 0 {
				continue
			}

			distance := math.Abs(float64(g1.X-g2.X)) + math.Abs(float64(g1.Y-g2.Y))
			result[key] = int(distance)
		}
	}

	sum := 0
	for _, v := range result {
		sum += v
	}

	fmt.Println(sum)
}

// 1: 4,0
// 2: 9,1
// 3: 0,2

// 1->2 9-4 + 1-0 == 6
// 1->3 0-4 + 2-0 == 6 //ABS
func findGalaxies(g Grid) []Coord {
	gals := []Coord{}
	g.Each(func(s string, x, y int) {
		if s == "#" {
			gals = append(gals, Coord{X: x, Y: y})
		}
	})
	return gals
}

func expand(g *Grid) {
	expandRows := []int{}
	for y := 0; y < g.Height(); y++ {
		row := g.GetRow(y)
		if !strings.Contains(row, "#") {
			expandRows = append(expandRows, y)
		}
	}
	fmt.Println(expandRows)

	dotRow := strings.Repeat(".", g.Width())
	expandedCount := 0
	for _, row := range expandRows {
		g.InsertRow(row+expandedCount, dotRow)
		expandedCount++
	}

	g.Print()

	expandCols := []int{}
	for x := 0; x < g.Width(); x++ {
		hasGalaxy := false
		for y := 0; y < g.Height(); y++ {
			loc := g.Get(x, y)
			if loc == "#" {
				hasGalaxy = true
			}
		}
		if !hasGalaxy {
			expandCols = append(expandCols, x)
		}
	}
	fmt.Println(expandCols)

	expandedCount = 0
	for _, col := range expandCols {
		for y := 0; y < g.Height(); y++ {
			g.InsertInRow(col+expandedCount, y, ".")
		}
		expandedCount++
	}
}
