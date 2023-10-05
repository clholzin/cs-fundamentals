package main

import (
	"math"
	"testing"
)

/*
Given two strings, str1 and str2, find the shortest substring in str1 such that str2 is a subsequence of that substring.

A substring is defined as a contiguous sequence of characters within a string. A subsequence is a sequence that can be derived from another sequence by deleting zero or more elements without changing the order of the remaining elements.

Let’s say you have the following two strings:

str1 = “abbcb”

str2 = “ac”

In this example, “
abbc
” is a substring of str1, from which we can derive str2 simply by deleting both the instances of the character
b
. Therefore, str2 is a subsequence of this substring. Since this substring is the shortest among all the substrings in which str2 is present as a subsequence, the function should return this substring, that is, “
abbc
”.

If there is no substring in str1 that covers all characters in str2, return an empty string.

If there are multiple minimum-length substrings that meet the subsequence requirement, return the one with the left-most starting index.
*/

func TestMinWindow(t *testing.T) {
	minWindow("", "")
}

func minWindow(str1, str2 string) string {
	minSubsequence := ""
	minSubLen := math.Inf(1)
	str1Len, str2Len := len(str1), len(str2)
	indexS1, indexS2 := 0, 0
	for indexS1 < str1Len {
		if str1[indexS1] == str2[indexS2] { // matched
			indexS2++ // increment sequence
			// if at the end of sequence
			if indexS2 == str2Len {
				start, end := indexS1, indexS1
				indexS2--
				for indexS2 >= 0 {
					if str1[start] == str2[indexS2] {
						indexS2--
					}
					start--
				}
				start++
				if float64(end-start+1) < minSubLen { // capture length of window and compare with previous minSubLen
					minSubLen = float64(end - start + 1)
					minSubsequence = str1[start : end+1]
				}
				indexS1 = start
				indexS2 = 0
			}
		}
		indexS1++
	}
	return minSubsequence
}
