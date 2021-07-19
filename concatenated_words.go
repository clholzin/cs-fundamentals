package fundamentals

/*
Time Complexity #
The time complexity of the above algorithm will be O(N * M * Len)O(N∗M∗Len) where ‘N’ is the number of characters in the given string, ‘M’ is the total number of words, and ‘Len’ is the length of a word.

Space Complexity #
The space complexity of the algorithm is O(M)O(M) since at most, we will be storing all the words in the two HashMaps. In the worst case, we also need O(N)O(N) space for the resulting list. So, the overall space complexity of the algorithm will be O(M+N).O(M+N).
*/

func findAllConcatenatedWords(st string, words []string) []int {
	wordFreq := make(map[string]int)
	for _, word := range words {
		wordFreq[word]++
	}

	resultsIndices := make([]int, 0)
	wordsCount := len(words)
	wordLength := len(words[0])

	for i := 0; i <= len(st)-wordsCount*wordLength; i++ {
		wordsSeen := make(map[string]int)
		for j := 0; j < wordsCount; j++ {
			nextWordIndex := i + j*wordLength
			word := st[nextWordIndex : nextWordIndex+wordLength]

			if _, ok := wordFreq[word]; !ok {
				break
			}

			wordsSeen[word]++

			if wordsSeen[word] > wordFreq[word] {
				break
			}

			if j+1 == wordsCount {
				resultsIndices = append(resultsIndices, i)
			}

		}
	}

	return resultsIndices
}
