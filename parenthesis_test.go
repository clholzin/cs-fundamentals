package fundamentals

import (
	"fmt"
	"testing"
)

func TestParts(t *testing.T) {

	fmt.Println(parenCount("()"))                 //1
	fmt.Println(parenCount("(())"))               //2
	fmt.Println(parenCount("((()()()))"))         //3
	fmt.Println(parenCount("((()()(())))"))       //4
	fmt.Println(parenCount("((((((((()))))))))")) //9
	fmt.Println(parenCount("((()))(())"))         //3 - fail 4
	fmt.Println(parenCount("(())((()))"))         //3 - fail 4

	fmt.Println(parenCheck("()"))                //true
	fmt.Println(parenCheck("(d(ds))"))           // true
	fmt.Println(parenCheck("(d(ds)"))            // false
	fmt.Println(parenCheck("(d(dass(()()())))")) //true
	fmt.Println(parenCheck("(d(dass(()(())))"))  //false
	fmt.Println(parenCheck("(()))"))             //false
	fmt.Println(parenCheck("(((((((())))))))"))  //true

}
