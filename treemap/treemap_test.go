package treemap

import (
	"math/rand"
	"testing"
	"time"
	"sort"
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

	if tm.Size() != 2 {
		t.Error()
	}

	tm.Remove("666")
	if tm.Size() != 1 {
		t.Error()
	}
}

func TestBalance(t *testing.T) {
	tm := New()
	for i := 0; i < 20000; i++ {
		k := getRandomStr(20)
		tm.Put(k, i)
		if i % 17 == 0 {
			tm.Remove(k)
		}
	}
	
	// Test the root is black.
	if tm.root.color != BLACK {
		t.Fatal("The root should be black")
	}
	// Test whether there is a continuous red nodes.
	checkRedNodes(tm.root, 0, t)
	// Test the number of black nodes per path is the same.
	checkBlackNodes(tm.root, t)
}

func checkRedNodes(x *Entry, num int, t *testing.T) {
	if x == nil {
		return
	}
	if x.color == RED {
		num++
	} else {
		num = 0
	}
	if num > 1 {
		t.Fatal("There is a continuous red nodes.")
	}
	checkRedNodes(x.left, num, t)
	checkRedNodes(x.right, num, t)
}

func checkBlackNodes(x *Entry, t *testing.T) (sum int) {
	if x == nil {
		return 1
	}
	leftSum := checkBlackNodes(x.left, t)
	rightSum := checkBlackNodes(x.right, t)
	if leftSum != rightSum {
		t.Logf("leftSum: %d, rightSum: %d", leftSum, rightSum)
		t.Fatal("The number of black nodes per path is the same.")
	}

	if x.color == BLACK {
		sum = leftSum + 1
	} else {
		sum = leftSum
	}
	return
}

func getRandomStr(length uint) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	ret := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := uint(0); i < length; i++ {
		ret = append(ret, bytes[r.Intn(len(bytes))])
	}
	return string(ret)
}

func TestIter(t *testing.T) {
	tm := New()
	for i := 0; i < 20000; i++ {
		tm.Put(i, i)
	}

	iter := tm.EntryIterator()
	i := 0
	for iter.HasNext() {
		e := iter.Next()
		
		if e.GetKey() != i || e.GetValue() != i {
			t.Fatalf("key: %x, value: %s\n", e.GetKey(), e.GetValue())
		}

		i++
	}
	if i != 20000 {
		t.Error()
	}
}

func TestOrder(t *testing.T) {
	tm := New()
	nums := []int { 13, 8, 17, 1, 11 ,15, 25, 6, 22, 27 }

	for _, num := range nums {
		tm.Put(num, nil)
	}
	
	iter := tm.EntryIterator()
	i := 0
	for iter.HasNext() {
		e := iter.Next()
		nums[i] = e.GetKey().(int)
		t.Log(e.GetKey())
		i++
	}
	if !sort.IntsAreSorted(nums) {
		t.Error("Key should be sorted.")
	}
}
