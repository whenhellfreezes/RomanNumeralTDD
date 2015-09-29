package roman

import (
	"testing"
)

func TestIsMultiple(t *testing.T) {
	one := romanSymbol{1, "I"}
	ten := romanSymbol{10, "X"}
	five := romanSymbol{5, "V"}
	if !one.isMultiple() {
		t.Error("One not a multiple")
		t.Fail()
	}
	if !ten.isMultiple() {
		t.Error("Ten not a multiple")
		t.Fail()
	}
	if five.isMultiple() {
		t.Error("Five counted as a multiple")
		t.Fail()
	}
}
