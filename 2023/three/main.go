package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Schematic []string

func (s Schematic) get(x, y int) string {
	if y > -1 && y < len(s) && x > -1 && x < len(s[y]) {
		return string(s[y][x])
	}

	return ""
}

var s Schematic

func main() {
	file, err := os.Open("3.input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		t := scanner.Text()
		s = append(s, t)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// for y := 0; y < 1; y++ {
	// 	for x := 0; x < len(s[y]); x++ {
	// 		fmt.Print(s.get(x, y))
	// 	}
	// 	fmt.Println("")
	// }

	// result := 0
	gearMap := map[string][]int{}
	// WALK THE GRID
	for y := 0; y < len(s); y++ {
		numString := ""
		for x := 0; x < len(s[y]); x++ {
			current := s.get(x, y)
			if isDigit(current) {
				numString += current
			}

			if numString != "" && (!isDigit(current) || x == len(s[y])-1) {
				// fmt.Printf("found %s", numString)
				// if search(x, y, len(numString)) {
				// 	num, err := strconv.Atoi(numString)
				// 	if err != nil {
				// 		panic(err)
				// 	}
				// 	result += num
				// } else {
				// 	fmt.Print(" -> NOT COUNTED\n")
				// }
				// numString = ""

				gear := searchGears(x, y, len(numString))
				if gear != "" {
					partNum, err := strconv.Atoi(numString)
					if err != nil {
						panic(err)
					}
					parts := gearMap[gear]
					parts = append(parts, partNum)
					gearMap[gear] = parts
				}
				numString = ""
			}
		}
	}

	result := 0
	for _, v := range gearMap {
		gearRatio := 1
		if len(v) == 2 {
			fmt.Println(v)
			for _, partNum := range v {
				gearRatio *= partNum
			}
			result += gearRatio
		}
	}
	fmt.Println(gearMap)
	fmt.Println(result)
}

func isDigit(s string) bool {
	if _, err := strconv.Atoi(s); err == nil {
		return true
	}

	return false
}

func searchGears(x, y, len int) string {
	fmt.Println(x, y)
	for j := y - 1; j < y+2; j++ {
		for i := x - (len + 1); i <= x; i++ {
			s := s.get(i, j)
			if s == "" {
				continue
			}

			if s == "*" {
				fmt.Printf("%d-%d\n", i, j)
				return fmt.Sprintf("%d-%d", i, j)
			}
		}
	}

	return ""
}

func search(x, y, len int) bool {
	for j := y - 1; j < y+2; j++ {
		for i := x - (len + 1); i <= x; i++ {
			s := s.get(i, j)
			if s == "" {
				continue
			}

			if s != "." && !isDigit(s) {
				fmt.Printf(" -> SAW %s COUNTED\n", s)
				return true
			}
		}
	}

	return false
}
