package main

import (
	u "aoc/util"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
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
			str += "*"
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

func (m Map) insertInRow(x, y int, s string) {
	if y > -1 && y < len(m) && x > -1 && x < len(m[y]) {
		str := m[y]
		newStr := str[:x] + s + str[x:]
		m[y] = newStr
	}
}

func (m Map) insertRow(y int, s string) Map {
	if y > -1 && y < len(m) {
		m = append(m[:y+1], m[y:]...)
		m[y] = s
		return m
	}

	return m
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

func (m Map) height() int {
	return len(m)
}

func (m Map) width() int {
	return len(m[0])
}

func (m Map) Print() {
	fmt.Println()
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			tile := m.get(x, y)
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
		}
		fmt.Println("")
	}
	fmt.Println()
	color.Unset()
}

var m Map

func main() {
	f, err := os.Open("inputs/10.input")
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

	m2 := make(Map, len(m))
	copy(m2, m)

	m2.Print()

	big, inserted := expand(m2)
	big.Print()

	traceStart := findStart(big, "S")
	trace(traceStart, true, big)
	big.Print()

	starts := findFillStarts(big)
	fill(starts, big)
	big.Print()

	count := 0
	allChars := "|-7FLJ."
	big.each(func(s string, x, y int) {
		// did we insert it?
		manMade := u.Any(inserted, Coord{x, y})
		if strings.Contains(allChars, s) && !manMade {
			count++
			big.set(x, y, "I")
		}
	})
	big.Print()
	fmt.Println(count)

	// part one
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
}

func findStart(m Map, target string) Coord {
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			tile := m.get(x, y)
			if tile == target {
				return Coord{x: x, y: y}
			}
		}
	}

	panic("NO START FOUND")
}

func findFillStarts(m Map) []Coord {
	result := []Coord{}
	for y := 0; y < len(m); y++ {
		if y == 0 || y == len(m)-1 {
			for x := 0; x < len(m[y]); x++ {
				tile := m.get(x, y)
				if tile != "$" {
					result = append(result, Coord{x: x, y: y})
				}
			}
		} else {
			tile := m.get(0, y)
			if tile != "$" {
				result = append(result, Coord{x: 0, y: y})
			}
			x := len(m[y]) - 1
			tile = m.get(x, y)
			if tile != "$" {
				result = append(result, Coord{x: x, y: y})
			}
		}
	}

	return result
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

func canMove(m Map, c Coord, prev Coord, dir string) (Coord, bool) {
	fromTile := m.getTile(c)
	maybe := move(c, dir)
	if maybe == prev {
		// fmt.Println("cannot go to a previously seen space")
		return maybe, false
	}
	tile := m.getTile(maybe)
	if tile == "" {
		// fmt.Println("cannot go out of bounds")
		return maybe, false
	}

	allowedFrom := strings.Contains(allowedFromMap[fromTile], dir)
	allowedTo := strings.Contains(allowedToMap[dir], tile)
	return maybe, (allowedFrom && allowedTo)
}

func canFill(m Map, c Coord, dir string) (next Coord, canMove bool) {
	maybe := move(c, dir)
	maybeTile := m.getTile(maybe)
	if maybeTile == "" || maybeTile == "$" {
		return maybe, false
	}

	return maybe, true
}

type Stack []Coord

func (s *Stack) Push(c Coord) {
	*s = append(*s, c)
}

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

func expand(m Map) (Map, []Coord) {
	inserted := []Coord{}
	// double the size of the map passed in
	r := NewMap(2*m.width(), 2*m.height())
	m.each(func(s string, x, y int) {
		r.set(2*x, 2*y, s)
	})

	// every other row every other column
	leftPipes := "FL-S"
	rightPipes := "-7JS"
	for y := 0; y < r.height(); y += 2 {
		for x := 1; x < r.width(); x += 2 {
			leftX := x - 1
			rightX := x + 1

			left := r.get(leftX, y)
			right := r.get(rightX, y)

			if strings.Contains(leftPipes, left) &&
				strings.Contains(rightPipes, right) {
				c := Coord{x, y}
				r.setTile(c, "-")
				inserted = append(inserted, c)
			} else {
				r.set(x, y, "*")
			}
		}
	}

	topPipes := "F7|S"
	bottomPipes := "JL|S"
	for y := 1; y < r.height(); y += 2 {
		for x := 0; x < r.width(); x += 2 {
			top := r.get(x, y-1)
			bottom := r.get(x, y+1)
			if strings.Contains(topPipes, top) &&
				strings.Contains(bottomPipes, bottom) {
				c := Coord{x, y}
				r.setTile(c, "|")
				inserted = append(inserted, c)
			} else {
				r.set(x, y, "*")
			}
		}
	}

	return r, inserted
}

func expandCols(m Map) {
	rightEdges := "|J7"
	leftEdges := "|LF"

	needsExpandedCols := []int{}
	for x := 0; x < len(m[0]); x++ {
		for y := 0; y < len(m); y++ {
			tile := m.get(x, y)
			if strings.Contains(rightEdges, tile) {
				right := m.get(x+1, y)
				if strings.Contains(leftEdges, right) {
					if !u.Any(needsExpandedCols, x) {
						needsExpandedCols = append(needsExpandedCols, x)
					}
				}
			}
		}
	}

	previouslyExpanded := 0
	for _, v := range needsExpandedCols {
		//expand it
		x := v + previouslyExpanded
		for y := 0; y < len(m); y++ {
			curr := m.get(x, y)
			right := m.get(x+1, y)
			if strings.Contains(rightEdges, curr) {
				m.insertInRow(x+1, y, "*")
			} else { //if I don't have a right edge
				if right == "-" || curr == "-" {
					m.insertInRow(x+1, y, "-")
				} else if curr == "." {
					m.insertInRow(x+1, y, "*")
				} else if strings.Contains(leftEdges, right) {
					m.insertInRow(x+1, y, "*")
				} else {
					m.insertInRow(x+1, y, "-")
				}
			}
		}
		previouslyExpanded++
	}
}

func expandRows(m Map) Map {
	bottomEdges := "-LJ"
	topEdges := "-7F"

	needsExpandedRows := []int{}
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			tile := m.get(x, y)
			if strings.Contains(bottomEdges, tile) {
				down := m.get(x, y+1)
				if strings.Contains(topEdges, down) {
					if !u.Any(needsExpandedRows, y) {
						needsExpandedRows = append(needsExpandedRows, y)
					}
				}
			}
		}
	}

	previouslyExpanded := 0
	for _, v := range needsExpandedRows {
		//expand it
		row := ""
		y := v + previouslyExpanded
		for x := 0; x < len(m[y]); x++ {
			curr := m.get(x, y)
			down := m.get(x, y+1)
			if strings.Contains(bottomEdges, curr) {
				row += "*"
			} else { //if I don't have a bottom edge
				if down == "|" || curr == "|" {
					row += "|"
				} else if curr == "." || curr == "*" {
					row += "*"
				} else {
					row += "|"
				}
			}
		}
		m = m.insertRow(y+1, row)
		previouslyExpanded++
	}

	return m
}

// fill outsides -- i think
func fill(starts []Coord, m Map) {
	seen := map[Coord]bool{}
	var curr Coord

	stack := Stack{}
	for _, start := range starts {
		stack.Push(start)
	}

	for {
		if stack.Count() == 0 {
			break
		}

		curr = stack.Pop()
		seen[curr] = true
		m.setTile(curr, "O")

		// fmt.Println("curr", curr)

		for _, dir := range dirs {
			next, canMove := canFill(m, curr, dir)
			if seen[next] {
				// fmt.Println("skipping", next, dir, "seen it")
				continue
			}
			if canMove {
				// fmt.Println("pushing", next, dir)
				stack.Push(next)
			} else {
				// fmt.Println(next, "not allowed", "to the", dir)
			}
		}
	}
}

func trace(start Coord, draw bool, m Map) map[Coord]bool {
	steps := 0
	res := map[Coord]bool{}
	var prev Coord
	curr := start
	// fmt.Println("starting at", start)

	limiter := false
	i := 0
	for {
		i++
		if limiter {
			if i == 4 {
				return res
			}
		}

		for _, dir := range dirs {
			next, allowed := canMove(m, curr, prev, dir)
			if allowed {
				// fmt.Println(dir)
				tile := m.getTile(curr)
				if tile != "S" && draw {
					m.setTile(curr, "$")
				}
				if next == start {
					// fmt.Println("found the start again")
					return res
				}

				res[curr] = true

				steps++
				prev = curr
				curr = next
				break
			} else {
				// fmt.Println("not allowed to go", dir, "to", next, "from", curr)
			}
		}
	}
}
