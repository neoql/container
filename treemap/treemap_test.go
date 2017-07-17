package treemap

import (
	"testing"
)

func TestPutAndGet(t *testing.T) {
	tm := New()
	
	tm.Put("888", "Hello")
	tm.Put("666", "Hi")

	if tm.Get("888") != "Hello" || tm.Get("666") != "Hi" {
		t.Error()
	}

	tm.Put("888", "World")
	if tm.Get("888") != "World" {
		t.Error()
	}
}
