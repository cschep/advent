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

	exRows, exCols := expand(*g)

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

			tg1 := translate(g1, exRows, exCols)
			tg2 := translate(g2, exRows, exCols)

			distance := math.Abs(float64(tg1.X-tg2.X)) + math.Abs(float64(tg1.Y-tg2.Y))
			result[key] = int(distance)
		}
	}

	sum := 0
	for _, v := range result {
		sum += v
	}

	fmt.Println(sum)
}

func translate(c Coord, exRows []int, exCols []int) Coord {
	rc := 0
	for _, er := range exRows {
		if c.Y > er {
			rc++
		}
	}

	cc := 0
	for _, ec := range exCols {
		if c.X > ec {
			cc++
		}
	}

	// TODO: I don't actually know why this works
	//  works "fine" at 1
	mult := 999999
	return Coord{X: c.X + (cc * mult), Y: c.Y + (rc * mult)}
}

func findGalaxies(g Grid) []Coord {
	gals := []Coord{}
	soFar := 1
	g.Each(func(s string, x, y int) {
		if s == "#" {
			gals = append(gals, Coord{X: x, Y: y})
			g.Set(x, y, fmt.Sprintf("%d", soFar))
			soFar++
		}
	})
	return gals
}

func expand(g Grid) ([]int, []int) {
	expandRows := []int{}
	for y := 0; y < g.Height(); y++ {
		row := g.GetRow(y)
		if !strings.Contains(row, "#") {
			expandRows = append(expandRows, y)
		}
	}

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

	return expandRows, expandCols
}
