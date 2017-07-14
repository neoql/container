package treemap

import (
	"testing"
)

func TestPutAndGet(t *testing.T) {
	tm := New()
	pre, err := tm.Put("Hello", "World")

	if pre != nil {
		t.Error("pre shuould be nil.")
	}

	value, err := tm.Get("Hello")
	if value != "World" || err != nil {
		t.Error(value)
	}

	pre, err = tm.Put("Hello", "TreeMap")
	if pre != "World" {
		t.Error(pre)
	}

	tm.Put("Hi", "World")
	value, err = tm.Get("Hi")
	if value != "World" {
		t.Error(value)
	}
}
