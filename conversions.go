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
	parsedSymbols := make([]romanSymbol, numberOfChars)
	var k int
	//Go through all the characters
	for k = 0; k < numberOfChars; k++ {
		v := splitS[k]
		//Determine the correct symbol
		newestSymbol, err := fromChar(v, romanDict)
		if err != nil {
			return 0, err
		}
		parsedSymbols[k] = newestSymbol
		//Either add or subtract the value of the symbol depending on whether the next symbol is larger or not
		if k > 0 {
			previousValue := parsedSymbols[k-1].value
			if newestSymbol.value > previousValue {
				total -= previousValue
			} else {
				total += previousValue
			}
		}
		if k == numberOfChars-1 {
			total += newestSymbol.value
		}
	}
	//The above doesn't assure that we had a legal sequence. We check use
	//our naive total and doulbe check against the convertToRoman
	//function to see if the sequence was properly formed.
	backofthebook := convertToRoman(total)
	if backofthebook != symbols {
		return 0, errors.New("Illegal Sequence")
	}
	return total, nil
}
