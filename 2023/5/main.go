package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Map struct {
	srcName  string
	destName string
	entries  []MapEntry
}

type MapEntry struct {
	destStart int
	srcStart  int
	length    int
}

func (me *MapEntry) destEnd() int {
	return me.destStart + me.length
}

func (me *MapEntry) srcEnd() int {
	return me.srcStart + me.length
}

var maps []Map
var seeds []string

func main() {
	file, err := os.Open("5.input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "seeds") {
			_, seedsStr, found := strings.Cut(line, ":")
			if !found {
				panic("NO SEEDS")
			}
			seeds = strings.Split(strings.Trim(seedsStr, " "), " ")
		}

		if strings.Contains(line, "to") {
			srcName, destName, found := strings.Cut(strings.Trim(line, " map:"), "-to-")
			if !found {
				panic("NO -to-")
			}
			m := Map{
				srcName:  srcName,
				destName: destName,
				entries:  []MapEntry{},
			}

			for scanner.Scan() {
				line := scanner.Text()
				if line == "" {
					break
				}
				var dest, src, l int
				fmt.Sscanf(line, "%d %d %d", &dest, &src, &l)
				m.entries = append(m.entries, MapEntry{
					srcStart:  src,
					destStart: dest,
					length:    l,
				})
			}

			maps = append(maps, m)
		}
	}

	var wg sync.WaitGroup
	for i := 0; i < len(seeds); i += 2 {
		seed, err := strconv.Atoi(seeds[i])
		if err != nil {
			panic(err)
		}

		seedLength, err := strconv.Atoi(seeds[i+1])
		if err != nil {
			panic(err)
		}

		go processPair(seed, seedLength, &wg)
	}

	wg.Wait()
}

func processPair(start int, length int, wg *sync.WaitGroup) int {
	wg.Add(1)
	lowest := math.MaxInt32
	for j := 0; j < length; j++ {
		to := "seed"
		location := start + j
		for to != "location" {
			// fmt.Printf("%d maps to ", location)
			to, location = trace(to, location)
			// fmt.Printf("%s at %d\n", to, location)
		}
		if location < lowest {
			lowest = location
			fmt.Println("new lowest", lowest)
		}
	}

	wg.Done()
	return lowest
}

func trace(from string, location int) (string, int) {
	var m Map
	for _, v := range maps {
		if v.srcName == from {
			m = v
		}
	}

	result := location
	for _, e := range m.entries {
		if location >= e.srcStart && location < e.srcEnd() {
			diff := location - e.srcStart
			result = e.destStart + diff
		}
	}

	// fmt.Println("tracing", from, m.destName, location)
	return m.destName, result
}
