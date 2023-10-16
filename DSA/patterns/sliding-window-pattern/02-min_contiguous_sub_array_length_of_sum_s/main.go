package main

import (
	"log"
	"math"
)

func main() {
	inputs := []int{1, 2, 3, 2, 4, 5, 6}

	s := 6

	answer := GetMinSubArrayLength(inputs, s)

	log.Printf("answer is => %v", answer)

}

func GetMinSubArrayLength(A []int, s int) int {
	start := 0
	end := 0
	currSum := 0
	min := int(math.Inf(+1))

	for end = 0; end < len(A); end++ {
		// grow our window to update the current sum
		currSum += A[end]
		// did our window.sum reached or exceeded the given sum ?
		for currSum >= s {
			// record the min length because we reached the required sum (and we maybe exceed it too so we got a min length)
			min = findMin((end - start + 1), min)
			// shrink our window
			// shrink the window
			currSum = currSum - A[start]
			start++
		}
	}
	return min
}

func findMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
