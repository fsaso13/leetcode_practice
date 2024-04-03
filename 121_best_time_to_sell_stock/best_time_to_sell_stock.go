package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	bestProfit := 0
	lhs := 0
	rhs := 1

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

	for {
		if lhs == len(prices)-1 {
			break
		}

		buyPrice := prices[lhs]
		// fmt.Printf("Buy Day[%d] -> %d\n", lhs, buyPrice)

		sellPrice := prices[rhs]
		// fmt.Printf("Sell Day[%d] -> %d\n", rhs, sellPrice)

		periodProfit := sellPrice - buyPrice
		// fmt.Printf("Profit: %d\n", periodProfit)

		if periodProfit <= 0 {
			lhs++
			rhs = lhs + 1
		} else {
			if periodProfit > bestProfit {
				bestProfit = periodProfit
			}
			if rhs < len(prices)-1 {
				rhs++
			} else {
				lhs++
				rhs = lhs + 1
			}

		}
	}
	return bestProfit
}

func main() {
	fmt.Println("Potato")

	// Input: prices = [7,1,5,3,6,4]
	// Output: 5
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))

	// Input: prices = [7,6,4,3,1]
	// Output: 0
	prices = []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices))

	// Input: prices = [1,2,4,2,5,7,2,4,9,0,9]
	// Output: 9
	prices = []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0, 9}
	fmt.Println(maxProfit(prices))

	// Input: prices = [2,4,1,11,7]
	// Output: 10
	prices = []int{2, 4, 1, 11, 7}
	fmt.Println(maxProfit(prices))

	// Input: prices = [3,2,6,5,0,3]
	// Output: 4
	prices = []int{3, 2, 6, 5, 0, 3}
	fmt.Println(maxProfit(prices))

}
