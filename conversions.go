package roman

import (
	"errors"
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
	var previousValue int
	for k, v := range splitS {
		newestSymbol, err := fromChar(v, romanDict)
		if err != nil {
			return 0, err
		}
		if k > 0 {
			if newestSymbol.value > previousValue {
				total -= previousValue
			} else {
				total += previousValue
			}
		}
		previousValue = newestSymbol.value
		if k == numberOfChars-1 {
			total += newestSymbol.value
		}
	}
	backofthebook := convertToRoman(total)
	if backofthebook != symbols {
		return 0, errors.New("Illegal Sequence")
	}
	return total, nil
}
