package main

import "fmt"

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("中文测试"))
	fmt.Println(game([]int{2, 2, 3}, []int{3, 2, 1}))
}

func game(guess []int, answer []int) int {
	count := 0
	for k, v := range guess {
		if v == answer[k] {
			count++
		}
	}
	return count
}
