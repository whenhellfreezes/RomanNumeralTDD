package roman

import (
	"testing"
)

func TestIsPowerOfTen(t *testing.T) {
	one := romanSymbol{1, "I"}
	ten := romanSymbol{10, "X"}
	five := romanSymbol{5, "V"}
	if !one.isPowerOfTen() {
		t.Error("One should be a power of ten")
		t.Fail()
	}
	if !ten.isPowerOfTen() {
		t.Error("Ten should be a power of ten")
		t.Fail()
	}
	if five.isPowerOfTen() {
		t.Error("Five counted as a power of ten")
		t.Fail()
	}
	if !romanThousand.isPowerOfTen() {
		t.Error("One thousand should be a power of ten")
		t.Fail()
	}
}

func TestLevelTest(t *testing.T) {
	if !(romanOne.getLevel() == 1) {
		t.Error("I is not showing as level 1")
		t.Fail()
	}
	if romanFive.getLevel() == 4 {
		t.Error("romanFive should not be level 4")
		t.Fail()
	}
	if romanFiveHundred.getLevel() != 6 {
		t.Error("D is not showing as level 6")
	}
}

func TestFromChar(t *testing.T) {
	fromD, err := fromChar("D")
	if err != nil {
		t.Error("From char gave an error")
		t.Fail()
	}
	if !romanFiveHundred.equals(fromD) {
		t.Error("FiveHundred could not be made from D")
		t.Fail()
	}
	fromX, err := fromChar("X")
	if err != nil {
		t.Error("From char gave an error")
		t.Fail()
	}
	if romanHundred.equals(fromX) {
		t.Error("Hundred should not be made from X")
		t.Fail()
	}
}
