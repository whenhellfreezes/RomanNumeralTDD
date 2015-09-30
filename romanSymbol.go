package roman

import (
	"errors"
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
	romanDict        = defaultRomanDictionary()
)

type romanDictionary []romanSymbol

func defaultRomanDictionary() romanDictionary {
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

func (r romanSymbol) equals(other romanSymbol) bool {
	if r.value == other.value && r.character == other.character {
		return true
	}
	return false
}

func (r romanSymbol) isPowerOfTen() bool {
	for i := 0; i < 4; i++ {
		if int(math.Pow10(i)) == r.value {
			return true
		}
	}
	return false
}

func fromChar(s string) (romanSymbol, error) {
	for _, v := range romanDict {
		if v.character == s {
			return v, nil
		}
	}
	return romanOne, errors.New("Invalid Character")
}

func fromValue(k int) (romanSymbol, error) {
	for _, v := range romanDict {
		if v.value == k {
			return v, nil
		}
	}
	return romanOne, errors.New("Invalid Value in fromValue")

}
