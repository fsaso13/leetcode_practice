package main

import (
	"fmt"
	"strconv"
)

func lock2string(lock []int) string {
	return strconv.FormatInt(int64((lock)[0]), 10) + strconv.FormatInt(int64((lock)[1]), 10) + strconv.FormatInt(int64((lock)[2]), 10) + strconv.FormatInt(int64((lock)[3]), 10)
}

func checkPermutation(deadendsMap map[string]bool, target []int, c0 int, c1 int, c2 int, c3 int, minMoves *int, minCounter *int, lock *[]int) {

	strLock := strconv.FormatInt(int64(c0), 10) + strconv.FormatInt(int64(c1), 10) + strconv.FormatInt(int64(c2), 10) + strconv.FormatInt(int64(c3), 10)
	strTarget := strconv.FormatInt(int64((target)[0]), 10) + strconv.FormatInt(int64((target)[1]), 10) + strconv.FormatInt(int64((target)[2]), 10) + strconv.FormatInt(int64((target)[3]), 10)

	fmt.Println("Lock: " + strLock)
	fmt.Println("Target: " + strTarget)

	if c0 == 9 {
		(*lock)[0] = 9
	}

	if c1 == 9 {
		(*lock)[1] = 9
	}

	if c2 == 9 {
		(*lock)[2] = 9
	}

	if c3 == 9 {
		(*lock)[3] = 9
	}

	if strTarget == strLock {
		if *minMoves > *minCounter {
			*minMoves = *minCounter
		}
		return
	}

	if target[0] < c0 && target[1] < c1 && target[2] < c2 && target[3] < c3 {
		return
	}

	_, ok := deadendsMap[strLock]
	if ok {
		return
	}

	// if (*lock)[0] >= 9 || (*lock)[1] >= 9 || (*lock)[2] >= 9 || (*lock)[3] >= 9 {
	// 	return
	// }

	// Get circular distance/direction
	mod := 1 // modifier
	if target[0] > 10-target[0] {
		mod *= -1
	}

	if (*lock)[0] < 9 {
		// if (*lock)[0] < target[0] {
		(*minCounter)++
		checkPermutation(deadendsMap, target, c0+mod, c1, c2, c3, minMoves, minCounter, lock)
		(*minCounter)--
		// }
		// (*lock)[0] -= mod
	}

	// check c1
	// Get circular distance/direction
	mod = 1 // modifier
	if target[1] > 10-target[1] {
		mod *= -1
	}

	if (*lock)[1] < 9 {

		// if (*lock)[1] < target[1] {
		(*minCounter)++
		checkPermutation(deadendsMap, target, c0, c1+mod, c2, c3, minMoves, minCounter, lock)
		(*minCounter)--
		// }
		// (*lock)[1] -= mod
	}

	// return
	// check c2
	// Get circular distance/direction
	mod = 1 // modifier
	if target[2] > 10-target[2] {
		mod *= -1
	}

	if (*lock)[2] < 9 {
		// if (*lock)[2] < target[2] {
		(*minCounter)++
		checkPermutation(deadendsMap, target, c0, c1, c2+mod, c3, minMoves, minCounter, lock)
		(*minCounter)--
		// }
		// (*lock)[2] -= mod
	}

	// check c3
	// Positive distance
	// Get circular distance/direction
	mod = 1 // modifier
	if target[3] > 10-target[3] {
		mod *= -1
	}

	if (*lock)[3] < 9 {
		// if (*lock)[3] < target[3] {
		(*minCounter)++
		checkPermutation(deadendsMap, target, c0, c1, c2, c3+mod, minMoves, minCounter, lock)
		(*minCounter)--
		// }
		// (*lock)[3] -= mod
	}
}

func openLock(deadends []string, target string) int {
	/**			c0 c1 c2 c3
	 * LOCK -> "0""0""0""0"
	 */

	// Target
	t0, _ := strconv.Atoi(string(target[0]))
	t1, _ := strconv.Atoi(string(target[1]))
	t2, _ := strconv.Atoi(string(target[2]))
	t3, _ := strconv.Atoi(string(target[3]))
	tarInt := []int{t0, t1, t2, t3}

	fmt.Printf("\"%d\" \"%d\" \"%d\" \"%d\"\n", t0, t1, t2, t3)

	// Lock
	lock := []int{0, 0, 0, 0}

	deadendsMap := map[string]bool{}
	for _, val := range deadends {
		deadendsMap[val] = true
	}

	minMoves := 9999
	minCounter := 0
	checkPermutation(deadendsMap, tarInt, 0, 0, 0, 0, &minMoves, &minCounter, &lock)
	return minMoves
}

func main() {
	fmt.Println("Potato")

	// Input: deadends = ["0201","0101","0102","1212","2002"], target = "0202"
	// Output: 6
	deadends := []string{"0201", "0101", "0102", "1212", "2002"}
	target := "0202"
	fmt.Println(openLock(deadends, target))
}
