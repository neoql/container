package treemap

import (
	"testing"
)

func TestCompareStr(t *testing.T) {
	a, b := "Tom", "Jack"
	ret, err := compare(a, b)

	if err != nil {
		t.Error(err)
	}

	if ret != 1 {
		t.Errorf("%s is bigger than %s.", a, b)
	}

	ret, err = compare(b, a)

	if err != nil {
		t.Error(err)
	}

	if ret != -1 {
		t.Errorf("%s is bigger than %s.", a, b)
	}

	ret, err = compare(a, a)

	if err != nil {
		t.Error(err)
	}

	if ret != 0 {
		t.Errorf("%s is as big as %s.", a, b)
	}

	ret, err = compare(10, a)
	if err == nil {
		t.Error("10 is wrong type.")
	}

	ret, err = compare(a, 10)
	if err == nil {
		t.Error("10 is wrong type.")
	}

	ret, err = compare(20, 10)
	if err == nil {
		t.Error("10 and 20 are wrong type.")
	}
}
