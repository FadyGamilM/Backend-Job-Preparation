package main

// the public interface for recieving a non-sorted array and returns a sorted array
func MergeSort(items []int) []int {
	// when we have an array of length = 1, its sorted, so return this one
	if len(items) < 2 {
		return items
	}

	// divide the array into two parts by getting the middle
	middle := len(items) / 2

	// sort the left part
	sortedLeft := MergeSort(items[:middle])

	// sort the right part
	sortedRight := MergeSort(items[middle:])

	// merge them togeather and return the merged array
	return merge(sortedLeft, sortedRight)
}

func merge(left, right []int) []int {
	// if there is a possibility of comparing the two arrays together
	// in otherwords (if there are items in both arrays so we can compare them and sort) .. so compare them
	result := make([]int, 0) // allocate memory for slice of ints with zero length, so i have to use append to add items or i will get an error
	// as long as we have items in both of them to compare
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			// the left is smaller so put it in the result and shift the left
			result = append(result, left[0])
			left = left[1:] // eliminate the 0'th item

		} else {
			// the right is equale the left or smaller so put it in the result and shift the right
			result = append(result, right[0])
			right = right[1:] // eliminate the 0'th item
		}
	}

	// now one of them is empty .. so the other one is already sorted so we can put it in the sorted result array with 100% sure that we got a sorted result
	for j := 0; j < len(left); j++ {
		result = append(result, left[j])
	}

	for k := 0; k < len(right); k++ {
		result = append(result, right[k])
	}

	return result
}
