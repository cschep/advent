package main

import (
	u "aoc/util"
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("2024/4/4.input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	g := u.NewGrid(0, 0)
	for scanner.Scan() {
		line := scanner.Text()
		g.PushRow(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// keep a whole list of 'em
	result := [][]u.Coord{}

	for y := 0; y < g.Height(); y++ {
		for x := 0; x < g.Width(); x++ {

			// if we find an X then we can start scanning
			if g.Get(x, y) == "X" {
				path := []u.Coord{{X: x, Y: y}}
				target := "M"

				for j := -1; j <= 1; j++ {
					for i := -1; i <= 1; i++ {
						// the idea here is to pick a direction
						// then loop through "walking" in that direction
						for k := 1; k < 4; k++ {
							newX := x + (i * k)
							newY := y + (j * k)

							// fmt.Println("checking", x, y, i, j, newX, newY)
							if g.Get(newX, newY) == target {
								c := u.Coord{X: newX, Y: newY}
								path = append(path, c)

								if target == "M" {
									// fmt.Println("found M at", newX, newY)
									target = "A"
								} else if target == "A" {
									// fmt.Println("found A at", newX, newY)
									target = "S"
								} else if target == "S" {
									// fmt.Println("found S at", newX, newY)
									result = append(result, path)
									// g.PrintWithCoords(path)
									path = []u.Coord{}
									target = "M"
								}
							} else {
								target = "M"
								break
							}
						}
					}
				}
			}
		}
	}

	// fmt.Println(g.Get(0, 0))
	// fmt.Println(g.Get(0, 1))
	// fmt.Println(g.Get(0, 2))
	// fmt.Println(g.Get(0, 3))

	fmt.Println(len(result))
}
