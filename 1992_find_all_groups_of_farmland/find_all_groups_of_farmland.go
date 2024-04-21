package main

import (
	"fmt"
)

func expandGroup(land *[][]int, x int, y int, groups *[][]int, visited *[300][300]int) {
	height := len(*land)
	width := len((*land)[0])

	top := y
	bot := y
	right := x
	left := x

	// TOP
	for {
		if top+1 >= height || (*land)[top+1][x] == 0 {
			break
		} else if (*land)[top+1][x] == 1 {
			top++
		}
	}

	// BOT
	for {
		if bot-1 < 0 || (*land)[bot-1][x] == 0 {
			break
		} else if (*land)[bot-1][x] == 1 {
			bot--
		}
	}

	// RIGHT
	for {
		if right+1 >= width || (*land)[y][right+1] == 0 {
			break
		} else if (*land)[y][right+1] == 1 {
			right++
		}
	}

	// Left
	for {
		if left-1 < 0 || (*land)[y][left-1] == 0 {
			break
		} else if (*land)[y][left-1] == 1 {
			left--
		}
	}

	// res := []int{left, bot, right, top}
	for i := bot; i <= top; i++ {
		for j := left; j <= right; j++ {
			(*visited)[i][j] = 1
		}
	}
	res := []int{bot, left, top, right}

	*groups = append(*groups, res)
}

func findFarmland(land [][]int) [][]int {

	groups := [][]int{}
	visited := [300][300]int{}

	for y, row := range land {
		for x, val := range row {
			if val == 1 && visited[y][x] == 0 {
				expandGroup(&land, x, y, &groups, &visited)
			}
		}
	}

	return groups
}

func main() {
	fmt.Println("Potato")

	// Input: land = [[1,0,0],[0,1,1],[0,1,1]]
	// Output: [[0,0,0,0],[1,1,2,2]]
	land := [][]int{{1, 0, 0}, {0, 1, 1}, {0, 1, 1}}
	fmt.Println(findFarmland(land))

	// Input: land = [[1,1],[1,1]]
	// Output: [[0,0,1,1]]
	land = [][]int{{1, 1}, {1, 1}}
	fmt.Println(findFarmland(land))

	// Input: land = [[0]]
	// Output: []
	land = [][]int{{0}}
	fmt.Println(findFarmland(land))

	// Input: land = [[1,1],[0,0]]
	// Output: [[0,0,1,0]]
	land = [][]int{{1, 1}, {0, 0}}
	fmt.Println(findFarmland(land))
}
