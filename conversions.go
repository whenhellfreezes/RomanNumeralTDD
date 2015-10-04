package roman

import (
	"errors"
	_ "math"
	"strings"
)

func convertToRoman(input int) string {
	collection := mixinRomanDictionary()
	max := 0
	var maxElement romanSymbol
	for _, v := range collection {
		if v.value > max && v.value <= input {
			max = v.value
			maxElement = v
		}
	}
	if max == 0 {
		return ""
	}
	return maxElement.character + convertToRoman(input-max)
}

func convertToNumeric(symbols string) (int, error) {
	splitS := strings.Split(symbols, "")
	numberOfChars := len(splitS)
	if numberOfChars == 0 {
		return 0, nil
	}
	total := 0
	parsedSymbols := make([]romanSymbol, numberOfChars)
	var k int
	for k = 0; k < numberOfChars; k++ {
		v := splitS[k]
		newestSymbol, err := fromChar(v, romanDict)
		if err != nil {
			return 0, err
		}
		parsedSymbols[k] = newestSymbol
	}
	for k, v := range parsedSymbols {
		if k+1 < numberOfChars {
			if v.value < parsedSymbols[k+1].value {
				total -= v.value
			} else {
				total += v.value
			}
		} else {
			total += v.value
		}
	}
	backofthebook := convertToRoman(total)
	if backofthebook != symbols {
		return 0, errors.New("Illegal Sequence")
	}
	return total, nil
}
