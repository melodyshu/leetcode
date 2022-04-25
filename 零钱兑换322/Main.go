package main

import "fmt"

func min(a int, b int) int {
	if a >= b {
		return b
	} else {
		return a
	}
}

func coinChange(coins []int, amount int) int {
	maxAmount := amount + 1
	var dp = make([]int, maxAmount)
	for i := 0; i < len(dp); i++ {
		dp[i] = maxAmount
	}
	//base case
	dp[0] = 0
	//状态遍历
	for i := 0; i < len(dp); i++ {
		for j := 0; j < len(coins); j++ {
			if i-coins[j] < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[i-coins[j]]+1)
		}
	}
	if dp[amount] == maxAmount {
		return -1
	}
	return dp[amount]
}

func main() {
	var coins = []int{1, 2, 5}
	num := coinChange(coins, 13)
	fmt.Println(num)
}
