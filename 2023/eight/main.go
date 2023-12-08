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
		fmt.Println(line)
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

	fmt.Println(traverse("AAA", 0, 1))
}

func traverse(key string, instructionIndex int, count int) int {
	// always loop through instructions
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

	if nextKey == "ZZZ" {
		return count
	}

	return traverse(nextKey, instructionIndex+1, count+1)
}
