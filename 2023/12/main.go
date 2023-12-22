package main

import (
	"aoc/util"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/log"
)

type Result struct {
	s     string
	count int
}

func main() {
	log.SetTimeFormat(time.TimeOnly)
	// log.SetLevel(log.DebugLevel)

	f, err := os.Open("inputs/12.input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var sum int
	// var wg sync.WaitGroup
	for scanner.Scan() {
		line := scanner.Text()
		s, counts, _ := strings.Cut(line, " ")

		cs := strings.Split(counts, ",")
		is := util.Atois(cs)

		//part 2!
		zs, zis := unfold(s, is)

		// waiting++
		// go search(s, is, resChan)
		// go func() {
		// 	wg.Add(1)
		// 	res := int64(search(s, is, make([]int, len(is)), 0, 0))
		// 	atomic.AddInt64(&sum, res)
		// 	wg.Done()
		// }()

		// new memo per line
		memo := map[[3]int]int{}
		// sum += search(zs, zis, make([]int, len(zis)), 0, 0, memo)
		// sum += search(s, is, make([]int, len(is)), 0, 0, memo)
		ans := search2(zs, zis, 0, memo)
		fmt.Println(s, ans)
		sum += ans
	}

	// wg.Wait()

	// for res := range resChan {
	// 	waiting--
	// 	log.Infof("%s finished with %d correct", res.s, res.count)
	// 	sum += res.count
	// 	if waiting == 0 {
	// 		close(resChan)
	// 	}
	// }

	log.Infof("puzzle result: %d", sum)
}

func unfold(s string, cs []int) (string, []int) {
	zs := s + "?"
	r := strings.Repeat(zs, 5)
	ns := r[:len(r)-1]

	rcs := []int{}
	for i := 0; i < 5; i++ {
		rcs = append(rcs, cs...)
	}

	return ns, rcs
}

// programming is amazing
// todo: actually understand this
func search2(s string, cs []int, currentCount int, memo map[[3]int]int) int {
	key := [3]int{len(s), len(cs), currentCount}
	if v, ok := memo[key]; ok {
		return v
	}

	if len(s) == 0 {
		if len(cs) == 0 && currentCount == 0 || len(cs) == 1 && currentCount == cs[0] {
			return 1
		} else {
			return 0
		}
	}

	sum := 0
	c := s[0]
	// fmt.Println(s, cs, string(c), currentCount)
	for _, v := range [2]byte{'#', '.'} {
		if c == v || c == '?' {
			if v == '#' {
				// if currentCount+1 <= cs[0] {
				sum += search2(s[1:], cs, currentCount+1, memo)
				// }
			} else if v == '.' {
				//if we finished a group, move to the next block
				//if it is correct!
				if len(cs) > 0 && currentCount == cs[0] {
					sum += search2(s[1:], cs[1:], 0, memo)
				} else if currentCount == 0 { //or if we're not counting yet just keep going
					sum += search2(s[1:], cs, 0, memo)
				}
			}
		}
	}

	memo[key] = sum
	return sum
}

// hideous
type Args struct {
	s    string
	cs   string
	ncs  string
	ncsi int
	i    int
}

func args(s string, ncs []int, ncsi int, i int) Args {
	sncs := ""
	for _, v := range ncs[ncsi:] {
		sncs += fmt.Sprint(v)
	}

	return Args{
		// s:    s,
		ncs:  sncs,
		ncsi: ncsi,
		i:    i,
	}
}

// sdI"M 3qwe <-- chimi
func search(s string, cs []int, ncs []int, ncsi int, i int, memo map[Args]int) int {
	key := args(s, ncs, ncsi, i)
	// fmt.Println(key)
	// if v, ok := memo[key]; ok {
	// 	fmt.Println("CACHE HIT", key, v)
	// 	for k, v := range memo {
	// 		fmt.Println(k, v)
	// 	}
	// 	return v
	// }
	if i == len(s) {
		log.Debug("found the end", ncs)
		if sliceEq(cs, ncs) {
			memo[key] = 1
			return 1
		}
		memo[key] = 0
		return 0
	}
	fmt.Println(s, string(s[i]), ncs, ncsi, i)

	if s[i] == '?' {
		hashStr := replaceAtIndex(s, '#', i)
		dotStr := replaceAtIndex(s, '.', i)

		ans := search(hashStr, cs, ncs, ncsi, i, memo) + search(dotStr, cs, ncs, ncsi, i, memo)
		memo[key] = ans
		return ans
	} else if s[i] == '.' {
		// if it's not the first i and the previous is a hash
		// then we finished a group
		if i != 0 && s[i-1] == '#' {
			// if we finished a group and it doesn't match the solution
			// then it never will and we should return 0
			if cs[ncsi] != ncs[ncsi] {
				memo[key] = 0
				return 0
			}

			//if we're still here keep going with the nsci increased by 1
			ans := search(s, cs, ncs, ncsi+1, i+1, memo)
			memo[key] = ans
			return ans
		}
		ans := search(s, cs, ncs, ncsi, i+1, memo)
		memo[key] = ans
		return ans

	} else if s[i] == '#' {
		// again if we are starting a group that isn't
		// represented in the original counts this won't be a solution
		if ncsi > len(ncs)-1 {
			memo[key] = 0
			return 0
		}
		// the second we are sure one of the groups doens't match
		// get on outta there
		if ncs[ncsi]+1 > cs[ncsi] {
			memo[key] = 0
			return 0
		}

		// typical go verbosity
		nncs := make([]int, len(ncs))
		copy(nncs, ncs)
		nncs[ncsi]++

		ans := search(s, cs, nncs, ncsi, i+1, memo)
		memo[key] = ans
		return ans
	}

	panic("nope")
}

func searchStack(s string, counts []int) int {
	type SoFar struct {
		s           string
		i           int
		counts      []int
		countsIndex int
	}
	result := 0
	stack := util.Stack[SoFar]{{s: s, i: 0, counts: make([]int, len(counts))}}

	i := 0
	for {
		i++
		if stack.Count() == 0 {
			break
		}
		// if i > 10 {
		// 	break
		// }

		// fmt.Println(stack)
		curr := stack.Pop()
		if curr.i == len(curr.s) {
			if sliceEq(curr.counts, counts) {
				result++
			}
			continue
		}

		switch curr.s[curr.i] {
		case '?':
			first := SoFar{s: curr.s, i: curr.i, countsIndex: curr.countsIndex}
			first.s = replaceAtIndex(curr.s, '#', curr.i)
			first.counts = make([]int, len(curr.counts))
			copy(first.counts, curr.counts)

			second := SoFar{s: curr.s, i: curr.i, countsIndex: curr.countsIndex}
			second.s = replaceAtIndex(curr.s, '.', curr.i)
			second.counts = make([]int, len(curr.counts))
			copy(second.counts, curr.counts)

			stack.Push(first)
			stack.Push(second)
		case '#':
			if curr.countsIndex > len(counts)-1 {
				continue
			}
			curr.counts[curr.countsIndex]++
			curr.i++
			stack.Push(curr)
		case '.':
			// just got out of a group
			if curr.i != 0 && curr.s[curr.i-1] == '#' {
				if curr.counts[curr.countsIndex] != counts[curr.countsIndex] {
					continue
				}
				curr.countsIndex++
			}
			curr.i++
			stack.Push(curr)
		}
	}

	log.Infof("%s got result %d", s, result)
	return result
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func groupCounts(springs string) []int {
	splitSprings := strings.Split(springs, ".")
	newCounts := []int{}
	for _, group := range splitSprings {
		if len(group) > 0 {
			newCounts = append(newCounts, len(group))
		}
	}

	return newCounts
}

func sliceEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
