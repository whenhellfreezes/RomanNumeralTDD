package roman

import (
	"math"
)

var (
	romanOne         = romanSymbol{1, "I"}
	romanFive        = romanSymbol{5, "V"}
	romanTen         = romanSymbol{10, "X"}
	romanFifty       = romanSymbol{50, "L"}
	romanHundred     = romanSymbol{100, "C"}
	romanFiveHundred = romanSymbol{500, "D"}
	romanThousand    = romanSymbol{1000, "M"}
	romanDict        = DefaultRomanDictionary()
)

type romanDictionary []romanSymbol

func (slice romanDictionary) Less(i, j int) bool {
	return slice[i].value < slice[j].value
}

func (slice romanDictionary) Len() int {
	return len(slice)
}

func (slice romanDictionary) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func DefaultRomanDictionary() romanDictionary {
	container := make([]romanSymbol, 7)
	container[0] = romanOne
	container[1] = romanFive
	container[2] = romanTen
	container[3] = romanFifty
	container[4] = romanHundred
	container[5] = romanFiveHundred
	container[6] = romanThousand
	return container
}

type romanSymbol struct {
	value     int
	character string
}

func (r romanSymbol) getLevel() int {
	for k, v := range romanDict {
		if v.value == r.value {
			return k + 1
		}
	}
	return -1
}

func (r romanSymbol) isPowerOfTen() bool {
	for i := 0; i < 4; i++ {
		if int(math.Pow10(i)) == r.value {
			return true
		}
	}
	return false
}
