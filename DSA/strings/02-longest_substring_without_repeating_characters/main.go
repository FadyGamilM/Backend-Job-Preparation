package main

import "log"

// Failed !! not covering every corner case .. but its good way of thinking btw
func lengthOfLongestSubstring_v2(s string) int {
	max := 0
	curr := 0
	store := make(map[string]bool)
	for _, c := range s {
		if store[string(c)] {
			if max > curr {
				curr = 1
			} else {
				max = curr
				curr = 1
			}
			// re-initialize the store data structure
			store = make(map[string]bool)
			store[string(c)] = true

		} else {
			curr++
			store[string(c)] = true
		}
	}

	if max > curr {
		curr = 1
	} else {
		max = curr
		curr = 1
	}

	return max
}

func lengthOfLongestSubstring(s string) int {
	max := 0
	curr := 0
	store := make(map[string]bool)
	for i, char := range s {
		store[string(char)] = true
		curr = 1 // initalize not incremetn
		if i > len(s)-1 {
			break
		}
		for _, innerChar := range s[i+1:] {
			if store[string(innerChar)] {
				store = make(map[string]bool)
				if max <= curr {
					max = curr
				}
				break
			} else {
				store[string(innerChar)] = true
				curr++
			}
		}
		if max <= curr {
			max = curr
		}
	}
	if max <= curr {
		max = curr
	}

	return max
}

func main() {
	log.Println(lengthOfLongestSubstring("abbsa"))
	log.Println(lengthOfLongestSubstring("aaaaaa"))
	log.Println(lengthOfLongestSubstring(""))
	log.Println(lengthOfLongestSubstring("dvdf"))
}
