package main

import (
	"math"

	"github.com/charmbracelet/log"
)

func main() {
	// log.Info(coinChange([]int{1, 2, 5}, 100))
	// log.Info(nSum(5))
	log.Info(climbStairs(10))
}

// simple DP example
func nSum(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + i
	}
	log.Info(dp)
	return dp[n]
}

// objective function: f(i) is the number of distinct ways to reach the i-th stair
// base cases:
// f(0) = 1
// f(1) = 1
//
// recurrence relation ?
// f(n) = f(n-1) + f(n-2)
//
// order of execution?
// -> bottom-up // because we're building up from the first stair?
//
// where to look for the answer?
// f(n)

func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// recursion and memoization -- would someone call this DP?
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	memo := map[int]int{}
	var inner func(int) int
	inner = func(amount int) int {
		if amount == 0 {
			return 0
		}
		if amount < 0 {
			return math.MaxInt
		}

		// if we already did it just return it
		if v, ok := memo[amount]; ok {
			return v
		}

		fewest := math.MaxInt

		for _, c := range coins {
			res := inner(amount - c)

			if res < math.MaxInt {
				fewest = min(fewest, res+1)
			}
		}

		memo[amount] = fewest
		return fewest
	}

	res := inner(amount)
	if res == math.MaxInt {
		return -1
	}
	return res
}
