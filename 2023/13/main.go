package main

import (
	"aoc/util"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("inputs/13.input.small")
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
		fmt.Println(line)
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
		g.Print()
		vmirror, err := checkVertical(g)
		if err != nil {
			fmt.Println(err)

			// rotate
			r := rotate(g)
			// check again e.g. "horizontal"
			hmirror, err := checkVertical(r)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("HORIZONTAL", hmirror)
				summary += (hmirror + 1) * 100
			}
		} else {
			// we found a vertical!
			fmt.Println(vmirror)
			summary += vmirror + 1
		}
	}

	fmt.Println("summary:", summary)
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
	// fmt.Println("checking", x1, x2)
	found := true
	for y := 0; y < g.Height(); y++ {
		if g.Get(x1, y) != g.Get(x2, y) {
			found = false
		}
	}
	// fmt.Println(found)
	return found
}

func checkVertical(g util.Grid) (int, error) {
	startx := -1
	for x := 0; x < g.Width(); x++ {
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
					fmt.Println("WINNER")
					return startx, nil
				}
			}
		}
	}

	return -1, fmt.Errorf("not vertical mirror found")
}
