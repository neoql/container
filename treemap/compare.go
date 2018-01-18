package treemap

import (
	"strings"
)

// Comparator can be compare.
type Comparator interface {
	CompareTo(o Comparator) int
}

// CprString is comparable string.
type CprString string

// CompareTo implements Comparator.
func (str CprString) CompareTo(o Comparator) int {
	a := string(str)
	b := string(o.(CprString))

	return strings.Compare(a, b)
}

// CprInt is comparable int.
type CprInt int

// CompareTo implements Comparator.
func (i CprInt) CompareTo(o Comparator) int {
	a := int(i)
	b := int(o.(CprInt))

	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}

func compare(a, b Comparator) (ret int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()

	ret = a.CompareTo(b)
	return
}

// CprFloat64 is comparable float64
type CprFloat64 float64

// CompareTo implements Comparator.
func (f CprFloat64) CompareTo(o Comparator) int {
	a := float64(f)
	b := float64(o.(CprFloat64))

	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}
