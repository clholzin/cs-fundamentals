package fundamentals

import (
	"fmt"
	"math"
	"testing"

	"github.com/yourbasic/bit"
)

func TestBits(t *testing.T) {
	var val Bits = 23
	t.Logf("%b - %d", val, val)
	var s Bits
	for i := Bits(0); i <= 4; i++ {
		if i < 2 || (i%2) == 0 {
			s += Get(val, i)
			t.Logf("%d (%b - %d)", i, s, s)
		}
	}
	t.Log()
	var g Bits
	val = 222
	for i := Bits(0); i <= 8; i++ {
		if i < 2 || (i%2) == 0 {
			g += Get(val, i)
			t.Logf("%d (%b - %d)", i, g, g)
		}
	}
	t.Logf("view %b", val)
	t.Logf("set 1 %b", Set(val, 1))
	val = Toggle(val, 2)
	t.Logf("toggle 2 %b - %d", val, val)

	t.Logf("clear 0 %b", Clear(val, 0))

}

func TestBitsTwo(t *testing.T) {
	var b Bits
	b = Set(b, F0)
	b = Toggle(b, F2)
	for i, flag := range []Bits{F0, F1, F2} {
		fmt.Println(i, Has(b, flag))
	}
}

func TestBitsArr(t *testing.T) {
	// Sieve of Eratosthenes
	const n = 50
	sieve := bit.New().AddRange(2, n)
	sqrtN := int(math.Sqrt(n))
	for p := 2; p <= sqrtN; p = sieve.Next(p) {
		for k := p * p; k < n; k += p {
			//t.Logf("%d %d\n", k, p)
			sieve.Delete(k)
		}
	}
	t.Log(sieve)
}
