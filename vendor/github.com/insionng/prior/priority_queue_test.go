package prior

import (
	"fmt"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	pq := NewPriorityQueue()
	n1 := NewNode("bootstrap", func() {
		fmt.Println("bootstrap")
	}, 2)
	n2 := NewNode("start", 2, 3)
	n3 := NewNode(3, "value", 4)
	pq.Push(n1)
	pq.Push(n2)
	pq.Push(n3)

	v := pq.Pop()
	if v.GetKey().(string) != "bootstrap" {
		t.Fatal()
	} else {
		if function, okay := v.GetValue().(func()); okay {
			function()
		}
	}
	v = pq.Pop()
	if v.GetKey().(string) != "start" {
		t.Fatal()
	}
	v = pq.Pop()
	if v.GetKey().(int) != 3 {
		t.Fatal()
	}
	v = pq.Pop()
	if v != nil {
		t.Fatal()
	}

	pq.Push(n1)
	pq.Push(n2)
	pq.Push(n3)

	pq.Remove(n2.GetIndex())
	pq.Remove(3)

	v = pq.Pop()
	if v.GetKey().(string) != "bootstrap" {
		t.Fatal()
	}
	v = pq.Pop()
	if v.GetKey().(int) != 3 {
		t.Fatal()
	} else {
		if v.GetValue().(string) != "value" {
			t.Fatal()
		}
	}
	v = pq.Pop()
	if v != nil {
		t.Fatal()
	}
}
