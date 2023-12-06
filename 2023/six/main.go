package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("6.input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	times := []int{}
	distances := []int{}

	scanner.Scan()
	line := scanner.Text()

	_, timeStr, _ := strings.Cut(line, ":")
	timeStrs := strings.Split(timeStr, " ")
	for _, v := range timeStrs {
		i, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		times = append(times, i)
	}

	scanner.Scan()
	line = scanner.Text()

	_, distanceStr, _ := strings.Cut(line, ":")
	distanceStrs := strings.Split(distanceStr, " ")
	for _, v := range distanceStrs {
		i, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		distances = append(distances, i)
	}

	fmt.Println(times, distances)

	result := 1
	for i, v := range times {
		wins := scoreRace(v, distances[i])
		// fmt.Println(v, wins)
		result *= wins
	}

	fmt.Println(result)
}

func scoreRace(time int, record int) int {
	wins := 0
	for i := 1; i < time; i++ {
		speed := i
		distance := (time - i) * speed
		if distance > record {
			wins++
		}
	}

	return wins
}
