package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Map []string

type Coord struct {
	x int
	y int
}

func NewMap(width, height int) Map {
	m := Map{}
	for y := 0; y < height; y++ {
		str := ""
		for x := 0; x < width; x++ {
			str += "."
		}
		m = append(m, str)
	}

	return m
}

func (m Map) get(x, y int) string {
	if y > -1 && y < len(m) && x > -1 && x < len(m[y]) {
		return string(m[y][x])
	}

	return ""
}

func (m Map) set(x, y int, s string) {
	if y > -1 && y < len(m) && x > -1 && x < len(m[y]) {
		asciiBytes := []byte(s)
		str := m[y]
		bytes := []byte(str)
		bytes[x] = asciiBytes[0]
		m[y] = string(bytes)
	}
}

func (m Map) getTile(c Coord) string {
	return m.get(c.x, c.y)
}

func (m Map) setTile(c Coord, tile string) {
	m.set(c.x, c.y, tile)
}

func (m Map) each(fn func(string, int, int)) {
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			tile := m.get(x, y)
			fn(tile, x, y)
		}
	}
}

func (m Map) Print() {
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			tile := m.get(x, y)
			fmt.Print(tile)
		}
		fmt.Println("")
	}
}

var m Map
var r Map

func main() {
	f, err := os.Open("inputs/10.input.small")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		m = append(m, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// var start Coord
	// for y := 0; y < len(m); y++ {
	// 	for x := 0; x < len(m[y]); x++ {
	// 		tile := m.get(x, y)
	// 		if tile == "S" {
	// 			// find and save start
	// 			start = Coord{x, y}
	// 		}
	// 	}
	// }
	//
	// r = NewMap(len(m[0]), len(m))
	//
	// loopSize := trace(start)
	// result := 0
	// if loopSize%2 == 0 {
	// 	result = loopSize / 2
	// } else {
	// 	result = loopSize/2 + 1
	// }
	//
	// r.print()
	// fmt.Println()
	// fmt.Println(result)

	// edges touching detection?
	// m.each(func(tile string, x, y int) {
	// 	if strings.Contains("|7J", tile) {
	// 		right := m.get(x+1, y)
	// 		if strings.Contains("|FLE", right) {
	// 			m.set(x, y, "E")
	// 			m.set(x+1, y, "E")
	// 		}
	// 	}
	//
	// 	if strings.Contains("-LJ", tile) {
	// 		down := m.get(x, y+1)
	// 		if strings.Contains("-F7E", down) {
	// 			m.set(x, y, "E")
	// 			m.set(x, y+1, "E")
	// 		}
	// 	}
	// })

	m.each(func(tile string, x, y int) {
		if tile == "." {
			//walk from here
		}
	})
	m.Print()

}

func move(c Coord, dir string) Coord {
	switch dir {
	case "n":
		return Coord{x: c.x, y: c.y - 1}
	case "e":
		return Coord{x: c.x + 1, y: c.y}
	case "s":
		return Coord{x: c.x, y: c.y + 1}
	case "w":
		return Coord{x: c.x - 1, y: c.y}
	}

	panic("tried to move in a bad direction")
}

var dirs = []string{"n", "e", "s", "w"}
var allowedToMap = map[string]string{
	"n": "|7FS",
	"e": "-J7S",
	"s": "|JLS",
	"w": "-FLS",
}
var allowedFromMap = map[string]string{
	"|": "ns",
	"-": "ew",
	"J": "wn",
	"F": "es",
	"L": "ne",
	"7": "ws",
	"S": "nesw",
}

func canMove(c Coord, prev Coord, dir string) (Coord, bool) {
	fromTile := m.getTile(c)
	maybe := move(c, dir)
	if maybe == prev {
		return maybe, false
	}
	tile := m.getTile(maybe)
	if tile == "" {
		return maybe, false
	}

	allowedFrom := strings.Contains(allowedFromMap[fromTile], dir)
	allowedTo := strings.Contains(allowedToMap[dir], tile)
	return maybe, (allowedFrom && allowedTo)
}

// is this all we need?
type Stack []Coord

func (s *Stack) Push(c Coord) {
	*s = append(*s, c)
}

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Pop() Coord {
	index := len(*s) - 1   // Get the index of the top most element.
	element := (*s)[index] // Index into the slice and obtain the element.
	*s = (*s)[:index]      // Remove it from the stack by slicing it off.
	return element
}

func (s *Stack) Count() int {
	return len(*s)
}

func trace(start Coord) int {
	steps := 0
	var prev Coord
	curr := start

	// i := 0
	for {
		// i++
		// if i == 20 {
		// 	break
		// }
		// fmt.Println(curr)

		tile := m.getTile(curr)
		r.setTile(curr, tile)
		for _, dir := range dirs {
			next, allowed := canMove(curr, prev, dir)
			if allowed {
				if next == start {
					return steps
				}

				steps++
				prev = curr
				curr = next
				break
			}
		}
	}
}
