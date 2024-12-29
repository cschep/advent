package main

import (
	"aoc/util"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("inputs/13.input")
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

	summary := 0
	for _, g := range grids {
		summary += checkAllSmudges(g)
	}

	fmt.Println("summary:", summary)
}

func checkAllSmudges(g util.Grid) int {
	// find the original reflection line
	orig, _ := check(g, CheckResult{})

	//part2 go through each one and flip it and check everything
	fmt.Println("CHECKING SMUDGES")
	g.Print()
	i := 0
	for y := 0; y < g.Height(); y++ {
		for x := 0; x < g.Width(); x++ {
			i++
			s := g.Get(x, y)
			new := "#"
			if s == "#" {
				new = "."
			}
			g.Set(x, y, new)
			res, found := check(g, orig)
			g.Set(x, y, s)

			if found {
				fmt.Println("FOUND", res)
			}

			if found && !res.eq(orig) {
				return res.amount
			}

			// if i == 1 {
			// 	return 0
			// }
		}
		fmt.Println()
	}

	panic("no solutions found for smudges")
}

type CheckResult struct {
	location  int
	amount    int
	direction string
}

func (cr CheckResult) eq(cr2 CheckResult) bool {
	return cr.location == cr2.location && cr.direction == cr2.direction && cr.amount == cr2.amount
}

//

func check(g util.Grid, ignore CheckResult) (CheckResult, bool) {
	ignoreLocation := -1
	if ignore.direction == "v" {
		ignoreLocation = ignore.location
	}
	vmirror, found := checkVertical(g, ignoreLocation)
	if found {
		return CheckResult{location: vmirror, amount: vmirror + 1, direction: "v"}, true
	}

	// check again rotated e.g. "horizontal"
	r := rotate(g)
	ignoreLocation = -1
	if ignore.direction == "h" {
		ignoreLocation = ignore.location
	}
	hmirror, found := checkVertical(r, ignoreLocation)
	if found {
		return CheckResult{location: hmirror, amount: (hmirror + 1) * 100, direction: "h"}, true
	}

	return CheckResult{}, false
}

// 90 deg counter clockwise
func rotate(g util.Grid) util.Grid {
	res := util.Grid{}
	for x := g.Width() - 1; x >= 0; x-- {
		rowStr := ""
		for y := 0; y < g.Height(); y++ {
			rowStr += g.Get(x, y)
		}
		res.PushRow(rowStr)
	}

	return res
}

func checkColumns(x1, x2 int, g util.Grid) bool {
	found := true
	for y := 0; y < g.Height(); y++ {
		if g.Get(x1, y) != g.Get(x2, y) {
			found = false
		}
	}
	return found
}

func checkVertical(g util.Grid, ignore int) (int, bool) {
	startx := -1
	for x := 0; x < g.Width(); x++ {
		if x == ignore {
			continue
		}
		if g.Get(x, 0) == g.Get(x+1, 0) {
			startx = x
			found := checkColumns(x, x+1, g)
			if found {
				// fmt.Println("MIRROR ON", x)
				x1 := x - 1
				x2 := x + 2
				for checkColumns(x1, x2, g) {
					x1 = x1 - 1
					x2 = x2 + 1
				}
				if x1 < 0 || x2 > g.Width()-1 {
					// fmt.Println("WINNER", startx)
					return startx, true
				}
			}
		}
	}

	return -1, false
}
