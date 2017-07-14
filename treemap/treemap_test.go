package treemap

import (
	"testing"
)

func TestPutAndGet(t *testing.T) {
	tm := New()
	pre, err := tm.Put("Hello", "World")

	t.Log(pre)
	t.Log(err)

	value, err := tm.Get("Hello")
	t.Log(err)
	if value != "World" || err != nil {
		t.Error(value)
	}
}
