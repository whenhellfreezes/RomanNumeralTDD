package roman

import (
	"testing"
)

func TestConvertToRoman(t *testing.T) {
	if convertToRoman(1006) != "MVI" {
		t.Error("Didn't get expected conversion for 1006")
		t.Fail()
	}

	if convertToRoman(94) != "XCIV" {
		t.Error("Didn't get expected conversion for 94")
		t.Fail()
	}
}

func TestConvertToNumber(t *testing.T) {
	answer, err := convertToNumeric("MVI")
	if err != nil || answer != 1006 {
		t.Error("Didn't get expected conversion for MVI")
		t.Fail()
	}
	answer, err = convertToNumeric("MVII")
	if err != nil || answer != 1007 {
		t.Error("Didn't get expected conversion for MVII")
		t.Fail()
	}
	answer, err = convertToNumeric("MCMM")
	if err == nil {
		t.Error("Expected an error for MCMM")
		t.Fail()
	}
	answer, err = convertToNumeric("XXIV")
	if err != nil || answer != 24 {
		t.Error("Didn't get expected conversion for XXIV")
		t.Fail()
	}
	answer, err = convertToNumeric("")
	if err != nil || answer != 0 {
		t.Error("Didn't get expected conversion for blank")
		t.Fail()
	}
}
