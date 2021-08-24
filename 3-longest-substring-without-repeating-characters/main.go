package main

import (
	"fmt"
	"math"
)

func main()  {
	fmt.Println(lengthOfLongestSubstring3("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring3("bbbbb"))
	fmt.Println(lengthOfLongestSubstring3("pwwkew"))
	fmt.Println(lengthOfLongestSubstring3(" "))
	fmt.Println(lengthOfLongestSubstring3("dvdf"))
	fmt.Println(lengthOfLongestSubstring3("abba"))
}

// Brute force
func lengthOfLongestSubstring(s string) int {
	strLen := len(s)
	maxLen := 0

	for i := 0; i < strLen; i++ {
		count := 0
		substr := make(map[uint8]int, 0)
		for k := i; k < strLen && strLen - i + 1 > maxLen; k++ {
			char := s[k]
			if _, exist := substr[char]; !exist {
				substr[char] = k
				count++
				if count > maxLen {
					maxLen = count
				}
				continue
			}
			break
		}
		if count > maxLen {
			maxLen = count
		}
	}

	return maxLen
}

// Sliding Window
func lengthOfLongestSubstring2(s string) int {
	strLen := len(s)
	maxLen := 0
	m := make(map[uint8]int)

	for left, right := 0, 0; right < strLen; right++ {
		char := s[right]
		if i, exists := m[char]; exists {
			left = int(math.Max(float64(i + 1), float64(left)))
		}
		m[char] = right

		maxLen = int(math.Max(float64(right - left + 1), float64(maxLen)))
	}

	return maxLen
}

// Optimized sliding window
func lengthOfLongestSubstring3(s string) int {
	strLen := len(s)
	maxLen := 0
	m := [128]int{}

	for left, right := 0, 0; right < strLen; right++ {
		char := s[right]
		if m[char] > 0 {
			left = int(math.Max(float64(m[char]), float64(left)))
		}
		m[char] = right + 1

		maxLen = int(math.Max(float64(right - left + 1), float64(maxLen)))
	}

	return maxLen
}
