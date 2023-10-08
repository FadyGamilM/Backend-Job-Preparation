package main

import "log"

func main() {
	nums := []int{-5, -3, -2, -1, -3}
	log.Printf("result is : %v", sortedSquares(nums))
}

func sortedSquares(nums []int) []int {
	// step (1) : square all numbers and this will achieve a complexity of O(n)
	for i, num := range nums {
		nums[i] = SquareNumbers(num)
	}

	log.Println(nums)

	// step (2) : apply the two pointer algorithm to sort the array in O(n) complexity
	// i will
	var ptr1 int
	var ptr2 int
	result := make([]int, len(nums))

	ptr1 = 0
	ptr2 = len(nums) - 1
	counter := 0
	for ptr1 < ptr2 {
		// if the items on the right are larger than the items on the left, thats good, move the right ptr to the left to sort the left part because the right item currently sorted (so far)
		if nums[ptr2] >= nums[ptr1] {
			result[counter] = nums[ptr2]
			counter++
			ptr2--
		} else if nums[ptr1] > nums[ptr2] {
			result[counter] = nums[ptr1]
			counter++
			ptr1++
		}
	}
	result[counter] = nums[ptr1]

	finalResult := make([]int, len(result))
	for i, item := range result {
		finalResult[len(result)-1-i] = item
	}

	return finalResult
}

func SquareNumbers(num int) int {
	return num * num
}
