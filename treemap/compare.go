package treemap

import (
	"strings"
	// "log"
)

// Comparator can be compare.
type Comparator interface {
	CompareTo(o interface{}) int
}

func compare(a, b interface{}) (ret int, err error) {
	defer func() {
		if e := recover(); e != nil {
			// log.Println(a, b, e)
			err = e.(error)
		}
	}()
	switch a.(type) {
	case string:
		ret = strings.Compare(a.(string), b.(string))
		return
	case int:
		x, y := a.(int), b.(int)
		if x < y {
			ret = -1
		} else if x > y {
			ret = 1
		} else {
			ret = 0
		}
		return
	default:
		ret = a.(Comparator).CompareTo(b)
		return
	}
}
