package fundamentals

type Bits uint8

const (
	F0 Bits = 1 << iota
	F1
	F2
)

func Get(b, flag Bits) Bits    { return (b & (1 << flag)) }
func Set(b, flag Bits) Bits    { return b | flag }
func Clear(b, flag Bits) Bits  { return b &^ flag }
func Toggle(b, flag Bits) Bits { return b ^ flag }
func Has(b, flag Bits) bool    { return b&flag != 0 }
