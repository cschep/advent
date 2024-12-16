package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func calibrationValue(line string) int {
	found := false
	first, last := 0, 0
	for _, ch := range line {
		i, err := strconv.Atoi(string(ch))
		if err != nil {
			continue
		}

		if !found {
			first = i
			found = true
		}
		last = i
	}
	return first*10 + last
}

func forward(line string) int {
	strs := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for i := 0; i < len(line); i++ {
		test := line[i:]
		if x, err := strconv.Atoi(test[0:1]); err == nil {
			return x
		}
		for j, v := range strs {
			if strings.HasPrefix(test, v) {
				return j + 1
			}
		}
	}

	panic("BAD INPUT")
}

func backwards(line string) int {
	strs := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for i := len(line); i >= 0; i-- {
		test := line[0:i]
		if x, err := strconv.Atoi(string(test[len(test)-1])); err == nil {
			return x
		}
		for j, v := range strs {
			if strings.HasSuffix(test, v) {
				return j + 1
			}
		}
	}

	panic("BAD INPUT")
}

func main() {
	file, err := os.Open("2023/1/1.input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// convert line from words to numbers
		first := forward(line)
		last := backwards(line)
		// fmt.Println(line, "first", first, "last", last)
		result += first*10 + last
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Println(result)
}

/*
go
python
javascript
typescript
ruby
crystal
swift
objective-c
c++
c
D
php
java
kotlin
scala
c#
vb.net
perl
bash ???????

//harder for me
Excel?
SQL
clojure
R
zig
nim
rust
scheme / racket
common lisp

//probably not
F#
haskell
*/
