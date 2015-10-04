package roman

import (
	"testing"
)

func TestFromChar(t *testing.T) {
	fromD, err := fromChar("D", romanDict)
	if err != nil {
		t.Error("From char gave an error")
		t.Fail()
	}
	if !romanFiveHundred.equals(fromD) {
		t.Error("FiveHundred could not be made from D")
		t.Fail()
	}
	fromX, err := fromChar("X", romanDict)
	if err != nil {
		t.Error("From char gave an error")
		t.Fail()
	}
	if romanHundred.equals(fromX) {
		t.Error("Hundred should not be made from X")
		t.Fail()
	}
}

func TestMixin(t *testing.T) {
	collection := mixinRomanDictionary()
	if len(collection) != 13 {
		t.Fail()
		t.Error("Failed mixins")
	}
}
