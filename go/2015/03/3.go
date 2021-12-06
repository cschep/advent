package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func part1() {
	type point struct {
		x int
		y int
	}

	f, err := os.Open("3.in")
	if err != nil {
		panic("file read failed")
	}
	defer f.Close()

	r := bufio.NewReader(f)
	currentPoint := point{0, 0}
	visited := make(map[point]bool)
	for {
		if c, _, err := r.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			visited[currentPoint] = true

			dir := string(c)

			switch dir {
			case "^":
				currentPoint.y += 1
			case ">":
				currentPoint.x += 1
			case "<":
				currentPoint.x -= 1
			case "v":
				currentPoint.y -= 1
			}
		}
	}

	log.Println(len(visited))
}

func part2() {
	type point struct {
		x int
		y int
	}

	f, err := os.Open("3.in")
	if err != nil {
		panic("file read failed")
	}
	defer f.Close()
	r := bufio.NewReader(f)

	santaCurrentPoint := point{0, 0}
	roboCurrentPoint := point{0, 0}
	var turnPoint = &santaCurrentPoint

	visited := make(map[point]bool)
	for {
		visited[*turnPoint] = true

		if c, _, err := r.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			dir := string(c)

			switch dir {
			case "^":
				(*turnPoint).y += 1
			case ">":
				(*turnPoint).x += 1
			case "<":
				(*turnPoint).x -= 1
			case "v":
				(*turnPoint).y -= 1
			}

			if turnPoint == &santaCurrentPoint {
				turnPoint = &roboCurrentPoint
			} else {
				turnPoint = &santaCurrentPoint
			}
		}
	}

	log.Println(len(visited))
}

func main() {
	part1()
	part2()
}
