// Package roman provides ...
package roman

var (
	romanOne         = romanSymbol{1, "I"}
	romanFive        = romanSymbol{5, "V"}
	romanTen         = romanSymbol{10, "X"}
	romanFifty       = romanSymbol{50, "L"}
	romanHundred     = romanSymbol{100, "C"}
	romanFiveHundred = romanSymbol{500, "D"}
	romanThousand    = romanSymbol{1000, "M"}
)

type romanSymbol struct {
	value     int
	character string
}

func (r romanSymbol) isMultiple() bool {
	if r.value%10 == 0 || r.value == 1 {
		return true
	}
	return false
}
