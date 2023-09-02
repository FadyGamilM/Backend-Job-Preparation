package main

import (
	"log"
	"math"
)

func main() {
	inputs := []int{1, 0, 3, 4, 1}
	k := 3
	answer := GetMaxAvgSubArray(inputs, k)
	log.Printf("answer is => %v", answer)
}

func GetMaxAvgSubArray(A []int, k int) float64 {
	start := 0
	end := 0
	max := float64(math.Inf(-1))
	avgSum := float64(0)
	for end = 0; end < len(A); end++ {
		// did we have a full window !
		if ((end - start) + 1) > k {
			// if the window end is at the end of the input array we can't expand our window anymore so we will compare the max with curr sum avg and return the result
			if (end + 1) > len(A) {
				// check the max value to be returned later ..
				max = math.Max(avgSum/float64(k), max)
				return max
			}
			avgSum = avgSum - float64(A[start]) + float64(A[end])
			max = math.Max(avgSum/float64(k), max)
			start += 1
		} else {
			// we still didn't reach the full window so we can keep adding the current item to our window and update the sum
			avgSum += float64(A[end])
			log.Println("the current max is : ", max, " where the start is : ", start, " and the end is : ", end)
		}
	}

	return max
}
