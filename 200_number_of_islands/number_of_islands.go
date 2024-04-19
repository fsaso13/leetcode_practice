package main

import (
	"fmt"
)

func islandExplorerHelper(grid *[][]byte, start_x int, start_y int) {

	// fmt.Printf("CELL[%d][%d]\n", start_x, start_y)
	height := len(*grid) - 1
	width := len((*grid)[0]) - 1

	// fmt.Printf("Height: %d\n", height)
	// fmt.Printf("Width: %d\n", width)

	top := start_y + 1
	bottom := start_y - 1
	right := start_x + 1
	left := start_x - 1

	// fmt.Printf("TOP: %d\n", top)
	// fmt.Printf("BOT: %d\n", bottom)
	// fmt.Printf("RIG: %d\n", right)
	// fmt.Printf("LEF: %d\n", left)
	// fmt.Println(string((*grid)[start_y][start_x]))

	// Mark cell as visited
	(*grid)[start_y][start_x] = '0'
	// Check neighbours if inside the grid
	// TOP Side
	if top <= height {
		// fmt.Println("Gone Top")
		if string((*grid)[top][start_x]) == "1" {
			// fmt.Println("Recursive")
			islandExplorerHelper(grid, start_x, top)
		}
	}

	// BOTTOM Side
	if bottom >= 0 {
		// fmt.Println("Gone Bot")
		if string((*grid)[bottom][start_x]) == "1" {
			// fmt.Println("Recursive")
			islandExplorerHelper(grid, start_x, bottom)
		}
	}
	// RIGHT Side
	if right <= width {
		// fmt.Println("Gone Right")
		if string((*grid)[start_y][right]) == "1" {
			// fmt.Println("Recursive")
			islandExplorerHelper(grid, right, start_y)
		}
	}

	// LEFT Side
	if left >= 0 {
		// fmt.Println("Gone Left")
		if string((*grid)[start_y][left]) == "1" {
			// fmt.Println("Recursive")
			islandExplorerHelper(grid, left, start_y)
		}
	}

	// fmt.Printf("Perimeter: %d\n", *perimeter)
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 { // Early return - Empty Input
		return 0
	} else if len(grid) == 1 && len(grid[0]) == 1 { // Early return - Single Cell
		if string(grid[0][0]) == "1" {
			return 1
		} else {
			return 0
		}
	}

	nIslands := 0
	start_x := -1
	start_y := -1

	// Find a plot of land
	for y, row := range grid {
		for x, val := range row {
			// fmt.Printf("Check Cell[%d][%d]: %d\n", x, y, val)
			if string(val) == "1" {
				islandExplorerHelper(&grid, x, y)
				nIslands++
			}
		}
		if start_x != -1 && start_y != -1 {
			break
		}
	}

	return nIslands
}

func main() {
	fmt.Println("Potato")

	// Input: grid = [
	// ["1","1","1","1","0"],
	// ["1","1","0","1","0"],
	// ["1","1","0","0","0"],
	// ["0","0","0","0","0"]
	// ]
	// Output: 1
	grid := [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}
	fmt.Println(numIslands(grid))

	// Input: grid = [
	// ["1","1","0","0","0"],
	// ["1","1","0","0","0"],
	// ["0","0","1","0","0"],
	// ["0","0","0","1","1"]
	// ]
	// Output: 3
	grid = [][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}
	fmt.Println(numIslands(grid))

	// Input: grid = [["1"]]
	// Output: 1
	grid = [][]byte{{'1'}}
	fmt.Println(numIslands(grid))

}
