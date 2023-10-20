package main

import (
	"log"
)

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

func lengthOfLongestSubstring_bruteforce(s string) int {
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

func lengthOfLongestSubstring(s string) int {
	curr, max := 0, 0
	store := make(map[rune]bool)
	window := make([]rune, 0)
	for _, char := range s {
		if store[char] {
			// check the last max we got before start shiftting my window
			if max < curr {
				max = curr
			}
			// start shiftting untile i got a new unique substring
			for store[char] {
				curr--
				shifttedChar := window[0]
				window = window[1:]
				// remove this char from the store hash table
				delete(store, shifttedChar)
			}
			// now add the character that made this confusion and increase the current length of the current unique substring
			curr++
			window = append(window, char)
			// add it to the hash table too
			store[char] = true
		} else {
			curr++
			window = append(window, char)
			store[char] = true
		}
	}
	if max < curr {
		max = curr
	}
	return max
}

func main() {
	log.Println(lengthOfLongestSubstring("abbsa"))
	log.Println(lengthOfLongestSubstring("aaaaaa"))
	log.Println(lengthOfLongestSubstring(""))
	log.Println(lengthOfLongestSubstring("dvdf"))
	log.Println(lengthOfLongestSubstring("pwwkew"))
}
