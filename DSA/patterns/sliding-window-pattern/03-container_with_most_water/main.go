package main

import (
	"log"
	"math"
)

func maxArea(height []int) int {
	// initialize 2 pointers
	p1 := 0
	p2 := len(height) - 1

	// initialize a max area value = 0 to start with
	max_area := float64(0)

	for p1 < p2 {
		// get the minimum height of them  * witdh
		curr_area := math.Min(float64(height[p1]), float64(height[p2])) * float64(p2-p1)
		max_area = math.Max(max_area, curr_area)
		if height[p1] < height[p2] {
			// shift p1 to the right
			p1++
		} else {
			// shift p2 to the left
			p2--
		}
	}

	return int(max_area)
}

func main() {
	log.Printf("result => %v \n", maxArea([]int{9, 7, 2, 8, 1, 6}))
}
