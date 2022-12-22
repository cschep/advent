package main

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func findWithPrefix(prefix string) {
	seed := "bgvyzdsv"
	i := 0
	for {
		nextInput := fmt.Sprintf("%s%d", seed, i)
		input := []byte(nextInput)
		res := fmt.Sprintf("%x\n", md5.Sum(input))
		if strings.HasPrefix(res, prefix) {
			fmt.Println("WOO!", i, res)
			break
		}
		i++
	}
}

func main() {
	findWithPrefix("00000")
	findWithPrefix("000000")
}
