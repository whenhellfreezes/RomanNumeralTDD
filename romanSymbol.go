package roman

import (
	"errors"
	"math"
)

func decompose(i int) (leading int, power int) {
	power = int(math.Log10(float64(i)))
	leading = int(float64(i) / math.Pow10(power))
	return
}

type romanSymbol struct {
	value     int
	character string
}

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

func fromChar(s string, dict romanDictionary) (romanSymbol, error) {
	for _, v := range dict {
		if v.character == s {
			return v, nil
		}
	}
	return romanOne, errors.New("Invalid Character")
}

func fromValue(k int, dict romanDictionary) (romanSymbol, error) {
	for _, v := range dict {
		if v.value == k {
			return v, nil
		}
	}
	return romanOne, errors.New("Invalid Value in fromValue")

}

func combine(first romanSymbol, second romanSymbol) romanSymbol {
	var result romanSymbol
	result.value = second.value - first.value
	result.character = first.character + second.character
	return result
}

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

func mixinRomanDictionary() romanDictionary {
	container := defaultRomanDictionary()
	result := defaultRomanDictionary()
	for _, v := range container {
		//Every roman symbol but "I" has a way to combine with another symbol
		if v.value != 1 {
			leading, power := decompose(v.value)
			//You can only have a 2 character symbol where the
			//prefix character is a power of ten. If the leading
			//digit of the higher symbol is 5 the two characters
			//are of the same power. If it's 1 then they are
			//different by a power.
			offset := 0
			if leading == 1 {
				offset = -1
			}
			//Create the appropriate prefix character.
			newSmybol, err := fromValue(int(math.Pow10(power+offset)), romanDict)
			if err != nil {
				panic("AGGGH")
			}
			//Add the 2 character symbol to the dictionary
			result = append(result, combine(newSmybol, v))
		}

	}
	return result
}
