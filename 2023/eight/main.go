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

	//
	starts := []string{}
	for k := range nodeMap {
		if strings.HasSuffix(k, "A") {
			starts = append(starts, k)
		}
	}

	// for {
	// 	select {
	// 	case <-stepChan:
	// 		traverse(starts[0], 0, 1, stepChan)
	// 		time.Sleep(time.Second * 5)
	// 	default:
	// 		println("Waiting for data")
	// 		time.Sleep(time.Duration(math.MaxInt64))
	// 	}
	// }

	chans := []chan string{}
	for _, start := range starts {
		ch := make(chan string)
		chans = append(chans, ch)
		go traverse(start, ch)
	}

	count := 1
	for {
		results := []string{}
		allZ := true
		for _, ch := range chans {
			next := <-ch
			fmt.Println(ch, next)
			if !strings.HasSuffix(next, "Z") {
				allZ = false
			}
			results = append(results, next)
		}
		if allZ {
			fmt.Println("WINNER", count)
			break
		}
		count++
	}
}

func traverse(key string, ch chan<- string) {
	instructionIndex := 0
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

		ch <- nextKey

		key = nextKey
		instructionIndex++
	}
}
