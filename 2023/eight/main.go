package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Node struct {
	left  string
	right string
}

var nodeMap map[string]Node
var instructions string

func main() {
	file, err := os.Open("8.input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	instructions = scanner.Text()

	scanner.Scan() // eat the newline

	nodeMap = map[string]Node{}
	for scanner.Scan() {
		line := scanner.Text()
		key, nodes, found := strings.Cut(line, " = ")
		if !found {
			panic("NOT FOUND")
		}

		left, right, found := strings.Cut(nodes, ", ")
		left = strings.Trim(left, "(")
		right = strings.Trim(right, ")")

		node := Node{
			left:  left,
			right: right,
		}
		nodeMap[key] = node
	}

	counts := []int{}
	for k := range nodeMap {
		if strings.HasSuffix(k, "A") {
			//find the counts for each one
			count := traverse(k)
			fmt.Println(k, count)
			counts = append(counts, count)
		}
	}

	// find the something common something
	a := counts[0]
	b := counts[1]
	fmt.Println(LCM(a, b, counts[2:]...))
}

func traverse(key string) int {
	instructionIndex := 0
	count := 1
	for {
		if instructionIndex > len(instructions)-1 {
			instructionIndex = 0
		}
		i := instructions[instructionIndex]

		var nextKey string
		if i == 'L' {
			nextKey = nodeMap[key].left
		} else {
			nextKey = nodeMap[key].right
		}

		if strings.HasSuffix(nextKey, "Z") {
			return count
		}

		key = nextKey
		instructionIndex++
		count++
	}
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
