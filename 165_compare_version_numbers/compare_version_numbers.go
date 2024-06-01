package main

import (
	"fmt"
	"strconv"
)

func compareVersion(version1 string, version2 string) int {
	lhs1 := 0
	rhs1 := 0
	lhs2 := 0
	rhs2 := 0

	got1 := false
	got2 := false
	for {
		// fmt.Printf("lhs1: %d , rhs1: %d\n", lhs1, rhs1)
		// fmt.Printf("lhs2: %d , rhs2: %d\n", lhs2, rhs2)

		if rhs1 >= len(version1) && rhs2 >= len(version2) {
			return 0
		}

		// fmt.Println(len(version1))
		if !got1 {
			if rhs1 >= len(version1)-1 {
				// fmt.Println("why")
				got1 = true
			} else if version1[rhs1] == '.' {
				rhs1--
				got1 = true
			} else {
				rhs1++
			}
		}

		// fmt.Println(len(version2))
		if !got2 {
			if rhs2 >= len(version2)-1 {
				// fmt.Println("why")
				got2 = true
			} else if version2[rhs2] == '.' {
				rhs2--
				got2 = true
			} else {
				rhs2++
			}
		}

		if got1 && got2 {
			lVer := int64(0)
			if rhs1 < len(version1) {
				lVer, _ = strconv.ParseInt(version1[lhs1:rhs1+1], 10, 64)
			}

			// fmt.Printf("lVer[%s]: %d\n", version1[lhs1:rhs1+1], lVer)
			// fmt.Printf("lhs1: %d , rhs1: %d\n", lhs1, rhs1)
			rVer := int64(0)
			if rhs2 < len(version2) {
				rVer, _ = strconv.ParseInt(version2[lhs2:rhs2+1], 10, 64)
			}
			// fmt.Printf("rVer[%s]: %d\n", version2[lhs2:rhs2+1], rVer)
			// fmt.Printf("lhs2: %d , rhs2: %d\n", lhs2, rhs2)

			if lVer < rVer {
				return -1
			} else if lVer > rVer {
				return 1
			} else {
				lhs1 = rhs1 + 2
				rhs1 = lhs1
				got1 = false
				lhs2 = rhs2 + 2
				rhs2 = lhs2
				got2 = false
			}
		}
	}
}

func main() {
	fmt.Println("Potato")

	// Input: version1 = "1.2", version2 = "1.10"
	// Output: -1
	v1 := "1.2"
	v2 := "1.10"
	fmt.Println(compareVersion(v1, v2))

	// Input: version1 = "1.01", version2 = "1.001"
	// Output: 0
	v1 = "1.01"
	v2 = "1.001"
	fmt.Println(compareVersion(v1, v2))

	// Input: version1 = "1.0", version2 = "1.0.0.0"
	// Output: 0
	v1 = "1.0"
	v2 = "1.0.0.0"
	fmt.Println(compareVersion(v1, v2))

	// Input: version1 = "1.00.1", version2 = "1.0.2"
	// Output: 0
	v1 = "1.00.1"
	v2 = "1.0.2"
	fmt.Println(compareVersion(v1, v2))
}
