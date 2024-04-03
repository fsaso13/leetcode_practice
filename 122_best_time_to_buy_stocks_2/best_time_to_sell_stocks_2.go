package main

import (
	"fmt"
)

func maxProfitHelper(prices []int, lhs int, rhs int) int {
	bestProfit := 0
	// lhs := 0
	// rhs := 1

	if len(prices) == 1 {
		return 0
	}

	if len(prices) == 2 {
		if prices[1]-prices[0] < 0 {
			return 0
		} else {
			return prices[1] - prices[0]
		}
	}

	// for {
	buyPrice := prices[lhs]
	sellPrice := prices[rhs]

	periodProfit := sellPrice - buyPrice

	if periodProfit <= 0 {
		lhs++
	} else {
		bestProfit += periodProfit
		maxProfitHelper(prices[rhs:])
	}
	// }
	return bestProfit
}

func maxProfit(prices []int) int {

}

func main() {
	fmt.Println("Potato")

	// Input: prices = [1,2,3,4,5]
	// Output: 5
	prices := []int{1, 2, 3, 4, 5}
	fmt.Println(maxProfitHelper(prices))
}
