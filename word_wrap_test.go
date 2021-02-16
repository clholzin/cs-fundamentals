package fundamentals

import (
	"testing"
)

func TestWordWrap(t *testing.T) {
	words := []string{"The", "day", "began",
		"as", "still", "as", "the", "night", "abrubtly",
		"lighted", "with", "brilliant", "flame"}
	expected := []string{
		"The-day-began",
		"as-still-as",
		"the-night",
		"abrubtly",
		"lighted-with",
		"brilliant",
		"flame",
	}
	maxlen := 13
	v := wraplines(words, maxlen)
	for i, d := range v {
		if expected[i] != d {
			t.Errorf("failed %s %s", expected[i], d)
		}
		t.Log(d)
	}
}
