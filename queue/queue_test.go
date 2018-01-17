package queue

import (
	"math/rand"
	"testing"
	"runtime"
	"sync"
)

func TestOneThread(t *testing.T) {
	q := New()
	ch := make(chan int, 8196)
	for i := 0; i < 8196; i++ {
		v := rand.Int()
		ch <- v
		q.Put(v)
	}

	if q.Size() != 8196 {
		t.Fatal("Not all values available")
	}

	close(ch)
	q.Close()
	for v := range ch {
		vv, _ := q.Pop()
		
		if v != vv.(int) {
			t.Fatal("Not all items received")
		}
	}
}

func TestMultiThread(t *testing.T) {
	runtime.GOMAXPROCS(2)

	ch := make(chan int, 8196)
	q  := New()

	g := sync.WaitGroup{}
	g.Add(2)

	go func() {
		for i := 0; i < 8196; i++ {
			v := rand.Int()
			ch <- v
			q.Put(v)
		}
		close(ch)
		q.Close()
		g.Done()
	}()

	go func() {
		for v := range ch {
			vv, _ := q.Pop()
			if v != vv.(int) {
				t.Fatal("Not all items received")
			}
		}
		g.Done()
	}()

	g.Wait()
}
