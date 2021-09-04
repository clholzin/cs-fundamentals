package fundamentals

import "strings"





func findOrderAlien(words []string) string {
   if len(words) == 0 {
      return ""
	 }

}

















// tags: topological sort

// time - V+N  where
// space - V+n

func findOrderAlien(words []string) string {

	if len(words) == 0 {
		return ""
	}

	edges := make(map[byte]int)
	parents := make(map[byte][]byte)

	for _, word := range words {
		for i := range word {
			edges[word[i]] = 0
			parents[word[i]] = make([]byte, 0)
		}
	}

	for i := 0; i < len(words)-1; i++ {
		w1, w2 := words[i], words[i+1]
		for j := 0; j < min(len(w1), len(w2)); j++ {
			parent, child := w1[j], w2[j]
			if parent != child {
				parents[parent] = append(parents[parent], child)
				edges[child] = edges[child] + 1
				break
			}
		}
	}

	queue := make([]byte, 0)
	for character, count := range edges {
		if count == 0 {
			queue = append(queue, character)
		}
	}

	sorted := strings.Builder{}

	for len(queue) > 0 {
		parent := queue[0]
		queue = queue[1:]
		sorted.WriteByte(parent)
		children := parents[parent]
		for _, child := range children {
			edges[child]--
			if edges[child] == 0 {
				queue = append(queue, child)
			}
		}
	}

	if sorted.Len() != len(edges) {
		return ""
	}

	return sorted.String()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
