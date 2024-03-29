package main

import "log"

// this version is for returning the numbers themselves
func twoSum_returns_numbers(nums []int, target int) []int {
	// base cases when we have an empty array or array of length = 1 => we have no result so we will return nil
	if len(nums) < 2 {
		return nil
	}

	// sort the array in o(nlogn) time complexity and o(n) space complexity
	nums = MergeSort(nums)

	log.Println("The sorted nums -> ", nums)

	// apply the two pointers pattern on the sorted nums
	p1 := 0
	p2 := len(nums) - 1
	found := false

	for p1 < p2 {
		if nums[p1]+nums[p2] == target {
			return []int{p1, p2}
		} else if nums[p1]+nums[p2] < target {
			p1++
		} else if nums[p1]+nums[p2] > target {
			p2--
		}
	}
	if !found {
		return nil
	}

	return nil
}

func twoSum(nums []int, target int) []int {
	// base cases when we have an empty array or array of length = 1 => we have no result so we will return nil
	if len(nums) < 2 {
		return nil
	}

	hashMap := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		_, ok := hashMap[nums[i]]
		if ok {
			return []int{hashMap[nums[i]], i}
		} else {
			// store the one which completes it
			hashMap[target-nums[i]] = i
		}
	}
	return nil
}

func main() {
	log.Println(twoSum([]int{2, 50, 30, 7}, 9))
}
