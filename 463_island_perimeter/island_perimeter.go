package main

import (
	"fmt"
)

func perimeterHelper(grid *[][]int, start_x int, start_y int, perimeter *int) {

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

	// Mark cell as visited
	(*grid)[start_y][start_x] = -1

	// fmt.Println(*grid)
	// Check neighbours if inside the grid
	// TOP Side
	if top <= height {
		// fmt.Println("Gone Top")
		if (*grid)[top][start_x] == 1 {
			// fmt.Println("Recursive")
			perimeterHelper(grid, start_x, top, perimeter)
		} else {
			if (*grid)[top][start_x] == 0 {
				// fmt.Println("Top +1")
				(*perimeter)++
			}
		}
	} else {
		// fmt.Println("Top +1 (border)")
		(*perimeter)++
	}

	// BOTTOM Side
	if bottom >= 0 {
		// fmt.Println("Gone Bot")
		if (*grid)[bottom][start_x] == 1 {
			// fmt.Println("Recursive")
			perimeterHelper(grid, start_x, bottom, perimeter)
		} else {
			if (*grid)[bottom][start_x] == 0 {
				// fmt.Println("Bot +1")
				(*perimeter)++
			}
		}
	} else {
		// fmt.Println("Bot +1 (border)")
		(*perimeter)++
	}

	// RIGHT Side
	if right <= width {
		// fmt.Println("Gone Right")
		if (*grid)[start_y][right] == 1 {
			// fmt.Println("Recursive")
			perimeterHelper(grid, right, start_y, perimeter)
		} else {
			if (*grid)[start_y][right] == 0 {
				// fmt.Println("Right +1")
				(*perimeter)++
			}
		}
	} else {
		// fmt.Println("Right +1(border)")
		(*perimeter)++
	}

	// LEFT Side
	if left >= 0 {
		// fmt.Println("Gone Left")
		if (*grid)[start_y][left] == 1 {
			// fmt.Println("Recursive")
			perimeterHelper(grid, left, start_y, perimeter)
		} else {
			if (*grid)[start_y][left] == 0 {
				// fmt.Println("Left +1")
				(*perimeter)++
			}
		}
	} else {
		// fmt.Println("Left +1(border)")
		(*perimeter)++
	}

	// fmt.Printf("Perimeter: %d\n", *perimeter)
}

func islandPerimeter(grid [][]int) int {
	/**
	 * 1 - Find a piece of the island
	 * 2 - Recursively check their neighbours
	 */

	if len(grid) == 0 { // Early return - Empty Input
		return 0
	} else if len(grid) == 1 && len(grid[0]) == 1 { // Early return - Single Cell
		if grid[0][0] == 1 {
			return 4
		} else {
			return 0
		}
	}

	perimeter := 0
	start_x := -1
	start_y := -1

	// Find a plot of land
	for y, row := range grid {
		for x, val := range row {
			// fmt.Printf("Check Cell[%d][%d]: %d\n", x, y, val)
			if val == 1 {
				start_x = x
				start_y = y
				break
			}
		}
		if start_x != -1 && start_y != -1 {
			break
		}
	}

	// If gris is just water, return
	if start_x == -1 && start_y == -1 {
		return 0
	}

	perimeterHelper(&grid, start_x, start_y, &perimeter)

	return perimeter
}

func main() {
	fmt.Println("potato")

	// Input: grid = [[0,1,0,0],[1,1,1,0],[0,1,0,0],[1,1,0,0]]
	// Output: 16
	grid := [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}
	fmt.Println(islandPerimeter(grid))

	// Input: grid = [[1]]
	// Output: 4
	grid = [][]int{{1}}
	fmt.Println(islandPerimeter(grid))

	// Input: grid = [[1,0]]
	// Output: 4
	grid = [][]int{{1, 0}}
	fmt.Println(islandPerimeter(grid))

	// Input: grid = [[0,1]]
	// Output: 4
	grid = [][]int{{0, 1}}
	fmt.Println(islandPerimeter(grid))
}
