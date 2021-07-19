package fundamentals

import (
	"fmt"
	"testing"
)

/*
  public static void main(String[] args) {
    List<Integer> result = WordConcatenation.findWordConcatenation("catfoxcat", new String[] { "cat", "fox" });
    System.out.println(result);
    result = WordConcatenation.findWordConcatenation("catcatfoxfox", new String[] { "cat", "fox" });
    System.out.println(result);
  }
*/

func TestConcatenatedWords(t *testing.T) {

	fmt.Println(findAllConcatenatedWords("catfoxcat", []string{"cat", "fox"}))

	fmt.Println(findAllConcatenatedWords("catcatfoxfox", []string{"cat", "fox"}))

}
