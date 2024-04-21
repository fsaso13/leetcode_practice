package main

import (
	"fmt"
)

func dfs(adjMatrix [][]int, source int, dest int, visited []int) bool {
	// fmt.Printf("source: %d dest: %d\n", source, dest)
	visited[source] = 1

	if source == dest {
		return true
	}

	// fmt.Println(adjMatrix[source])
	for _, neighbor := range adjMatrix[source] {
		// fmt.Println(visited)
		if visited[neighbor] != 1 {
			if dfs(adjMatrix, neighbor, dest, visited) {
				return true
			}
		}
	}

	return false
}

func validPath(n int, edges [][]int, source int, destination int) bool {
	adjMatrix := make([][]int, n)

	for _, links := range edges {
		from := links[0]
		to := links[1]

		adjMatrix[from] = append(adjMatrix[from], to)
		adjMatrix[to] = append(adjMatrix[to], from)
		// adjMatrix[to][from] = 1
	}

	// fmt.Println(adjMatrix)
	visited := make([]int, n)
	found := dfs(adjMatrix, source, destination, visited)
	// fmt.Println(visited)
	return found
}

func main() {
	fmt.Println("Potato")

	// Input: n = 3, edges = [[0,1],[1,2],[2,0]], source = 0, destination = 2
	// Output: true
	n := 3
	edges := [][]int{{0, 1}, {1, 2}, {2, 0}}
	source := 0
	dest := 2
	fmt.Println(validPath(n, edges, source, dest))

	// // Input: n = 6, edges = [[0,1],[0,2],[3,5],[5,4],[4,3]], source = 0, destination = 5
	// // Output: false
	n = 6
	edges = [][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}}
	source = 0
	dest = 5
	fmt.Println(validPath(n, edges, source, dest))

	// Input: n = 10, edges = [[0,7],[0,8],[6,1],[2,0],[0,4],[5,8],[4,7],[1,3],[3,5],[6,5]], source = 7, destination = 5
	// Output: false
	n = 10
	edges = [][]int{{0, 7}, {0, 8}, {6, 1}, {2, 0}, {0, 4}, {5, 8}, {4, 7}, {1, 3}, {3, 5}, {6, 5}}
	source = 7
	dest = 5
	fmt.Println(validPath(n, edges, source, dest))
}
