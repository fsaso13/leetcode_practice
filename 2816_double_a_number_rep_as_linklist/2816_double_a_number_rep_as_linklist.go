package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func buildList(numArr []int) *ListNode {
	// fmt.Println(numArr)
	head := ListNode{numArr[0], nil}

	prev := &head
	for ind, val := range numArr {
		if ind == 0 {
			continue
		}

		aux := ListNode{val, nil}
		prev.Next = &aux
		prev = &aux
	}

	return &head
}

func doubleIt(head *ListNode) *ListNode {
	numAsArray := []int{0}

	currNode := head
	pos := 0
	for currNode != nil {
		// fmt.Printf("val: %d\n", currNode.Val)
		numAsArray = append(numAsArray, currNode.Val)
		pos++
		currNode = currNode.Next
	}

	extra := 0
	sol := make([]int, len(numAsArray))
	for i := len(numAsArray) - 1; i >= 0; i-- {
		currNum := numAsArray[i]

		double := currNum * 2

		res := (double % 10) + extra
		sol[i] = res
		extra = double / 10
		// fmt.Println(res)
		// fmt.Println(extra)
	}
	// fmt.Println("pots")
	// fmt.Println(numAsArray)
	// fmt.Println(sol)
	if sol[0] == 0 {
		sol = sol[1:]
	}
	return buildList(sol)
}

func main() {
	fmt.Println("Potato")

	// Input: head = [1,8,9]
	// Output: [3,7,8]
	in := []int{1, 8, 9}
	head := buildList(in)
	fmt.Println(doubleIt(head))

	// Input: head = [1,8,9]
	// Output: [3,7,8]
	in = []int{9, 9, 9}
	head = buildList(in)
	fmt.Println(doubleIt(head))
}
