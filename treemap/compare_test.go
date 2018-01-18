package treemap

import (
	"testing"
)

func TestCompareStr(t *testing.T) {
	a, b := "Tom", "Jack"
	ret, err := compare(CprString(a), CprString(b))

	if err != nil {
		t.Error(err)
	}

	if ret != 1 {
		t.Errorf("%s is bigger than %s.", a, b)
	}

	ret, err = compare(CprString(b), CprString(a))

	if err != nil {
		t.Error(err)
	}

	if ret != -1 {
		t.Errorf("%s is bigger than %s.", a, b)
	}

	ret, err = compare(CprString(a), CprString(a))

	if err != nil {
		t.Error(err)
	}

	if ret != 0 {
		t.Errorf("%s is as big as %s.", a, b)
	}

	ret, err = compare(CprFloat64(10.0), CprString(a))
	if err == nil {
		t.Error("10 is wrong type.")
	}

	ret, err = compare(CprString(a), CprFloat64(10.0))
	if err == nil {
		t.Error("10 is wrong type.")
	}
}
