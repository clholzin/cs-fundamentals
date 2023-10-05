package main

import (
	"fmt"
	"strings"
	"testing"
)

/*
Given a string, s, of lowercase English characters and an integer, k, return the length of the longest substring after replacing at most k characters with any other lowercase English character so that all the characters in the substring are the same.
*/

func TestLongestRepeating(t *testing.T) {
	inputStrings := []string{"abcdefghkkkkkkkabcdefghkkkkkkkk", "aabccbb", "abbcb", "abccde", "abbcab", "bbbbbbbbb"}
	k := []int{8, 2, 1, 1, 2, 4}

	for i := 0; i < len(inputStrings); i++ {
		fmt.Printf("%d.\tInput String: '%s'\n", i+1, inputStrings[i])
		fmt.Printf("\tk: %d\n", k[i])
		fmt.Printf("\tLength of the longest substring with repeating characters: %d\n", longestrepeating(inputStrings[i], k[i]))
		fmt.Println(strings.Repeat("-", 100))
	}
}

func longestrepeating(s string, k int) int {
	stringLength := len(s)
	lengthOfMaxSubstring := 0
	start := 0
	charFreq := make(map[byte]int)
	mostFreqChar := 0
	for end := 0; end < stringLength; end++ {
		if _, ok := charFreq[s[end]]; !ok {
			charFreq[s[end]] = 1
		} else {
			charFreq[s[end]]++
		}

		mostFreqChar = max(mostFreqChar, charFreq[s[end]])
		if end-start+1-mostFreqChar > k {
			charFreq[s[start]]--
			start++
		}
		lengthOfMaxSubstring = max(end-start+1, lengthOfMaxSubstring)
		fmt.Println(string(s[end]), start, mostFreqChar, lengthOfMaxSubstring)

	}
	return lengthOfMaxSubstring
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
