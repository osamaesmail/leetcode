package main

import "fmt"

func main()  {
	//fmt.Println(longestPalindrome("aba"))
	//fmt.Println(longestPalindrome("a"))
	//fmt.Println(longestPalindrome("ac"))
	//fmt.Println(longestPalindrome("aa"))
	//fmt.Println(longestPalindrome("aacabdkacaa"))
	//fmt.Println(longestPalindrome2("abccccba"))
}

// dynamic programing
func longestPalindrome(s string) string {
	strLen := len(s)
	if strLen < 2 {
		return s
	}
	table := make([][]bool, strLen)
	strMax := ""
	lenMax := 0
	for i := range table{
		table[i] = make([]bool, strLen)
	}
	for i := 0; i < strLen; i++ {
		for j := 0; j < strLen; j++ {
			start, end := j, j + i
			isInnerOK := false
			if end >= strLen {
				break
			}
			if end - start < 2 || table[start + 1][end - 1] {
				isInnerOK = true
			}
			if isInnerOK && s[start] == s[end] {
				table[start][end] = true
				currLen := end - start + 1
				if currLen > lenMax {
					lenMax = currLen
					strMax = s[start:end+1]
				}
			}
			fmt.Println(start)
			fmt.Println(end)
			fmt.Println(isInnerOK)
			fmt.Println(table[start][end])
			fmt.Println(strMax)
			fmt.Println("-----------")
		}
	}
	return strMax
}

// expand around center
func longestPalindrome2(s string) string {
	strLen := len(s)
	if strLen < 2 {
		return s
	}
	lStart, maxLen := 0, 1
	for i := 0; i < strLen; {
		if strLen - i <= maxLen/2 {
			break
		}
		// skip duplicates
		k, j := i, i
		for k < strLen - 1 && s[k] == s[k+1] {
			k++
		}
		// expand around centre
		i = k + 1
		for k < strLen - 1 && j > 0 && s[k + 1] == s[j - 1] {
			k++
			j--
		}
		newLen := k - j + 1
		if newLen > maxLen {
			maxLen = newLen
			lStart = j
		}
	}
	return s[lStart:lStart + maxLen]
}