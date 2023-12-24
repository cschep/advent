package util

import (
	"fmt"

	"github.com/fatih/color"
)

type Grid []string

type Coord struct {
	X int
	Y int
}

func NewGrid(width, height int) *Grid {
	g := Grid{}
	for y := 0; y < height; y++ {
		str := ""
		for x := 0; x < width; x++ {
			str += "*"
		}
		g = append(g, str)
	}

	return &g
}

func (g Grid) Get(x, y int) string {
	if y > -1 && y < g.Height() && x > -1 && x < g.Width() {
		return string(g[y][x])
	}

	return ""
}

func (g *Grid) Set(x, y int, s string) {
	if y > -1 && y < g.Height() && x > -1 && x < g.Width() {
		asciiBytes := []byte(s)
		str := (*g)[y]
		bytes := []byte(str)
		bytes[x] = asciiBytes[0]
		(*g)[y] = string(bytes)
	}
}

func (g *Grid) InsertInRow(x, y int, s string) {
	if y > -1 && y < g.Height() && x > -1 && x < g.Width() {
		str := (*g)[y]
		newStr := str[:x] + s + str[x:]
		(*g)[y] = newStr
	}
}

func (g *Grid) InsertRow(y int, s string) {
	if y > -1 && y < g.Width() {
		*g = append((*g)[:y+1], (*g)[y:]...)
		(*g)[y] = s
	}
}

func (g *Grid) PushRow(s string) {
	*g = append(*g, s)
}

func (g *Grid) GetRow(y int) string {
	return (*g)[y]
}

func (g Grid) getTile(c Coord) string {
	return g.Get(c.X, c.Y)
}

func (g *Grid) setTile(c Coord, tile string) {
	g.Set(c.X, c.Y, tile)
}

func (g Grid) Each(fn func(s string, x int, y int)) {
	for y := 0; y < g.Height(); y++ {
		for x := 0; x < g.Width(); x++ {
			tile := g.Get(x, y)
			fn(tile, x, y)
		}
	}
}

func (g Grid) Height() int {
	return len(g)
}

func (g Grid) Width() int {
	if len(g) > 0 {
		return len(g[0])
	} else {
		return 0
	}
}

func (g Grid) rotateCC() Grid {
	res := Grid{}
	for x := g.Width() - 1; x >= 0; x-- {
		rowStr := ""
		for y := 0; y < g.Height(); y++ {
			rowStr += g.Get(x, y)
		}
		res.PushRow(rowStr)
	}

	return res
}

func (g Grid) rotateCW() Grid {
	res := Grid{}
	for x := 0; x < g.Width(); x++ {
		rowStr := ""
		for y := g.Height() - 1; y >= 0; y-- {
			rowStr += g.Get(x, y)
		}
		res.PushRow(rowStr)
	}

	return res
}

func (g *Grid) Print() {
	fmt.Println()
	fmt.Println("             111")
	fmt.Println("   0123456789012")
	for y := 0; y < g.Height(); y++ {
		for x := 0; x < g.Width(); x++ {
			if x == 0 {
				fmt.Printf("%.2d ", y)
			}
			tile := g.Get(x, y)
			if tile == "$" {
				color.Set(color.FgGreen)
			} else if tile == "I" {
				color.Set(color.FgRed)
			} else if tile == "O" {
				color.Set(color.FgCyan)
			} else if tile == "S" {
				color.Set(color.FgYellow)
			} else {
				color.Unset()
			}
			fmt.Print(tile)
			color.Unset()
		}
		fmt.Println("")
	}
	fmt.Println()
}
