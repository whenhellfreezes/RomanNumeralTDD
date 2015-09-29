// Package roman provides ...
package roman

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
