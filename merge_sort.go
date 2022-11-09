package main

import (
	"github.com/prog-1/list"
)

// Merge combines two sorted lists into a single sorted list.
func Merge(a, b *list.List) *list.List {
	curr1 := a.Front()
	curr2 := b.Front()
	c := new(list.List)
	for curr1 != nil && curr2 != nil {
		if curr1.Value > curr2.Value {
			c.PushBack(curr2.Value)
			curr2 = curr2.Next()
		} else {
			c.PushBack(curr1.Value)
			curr1 = curr1.Next()
		}
	}
	for curr1 != nil {
		c.PushBack(curr1.Value)
		curr1 = curr1.Next()
	}
	for curr2 != nil {
		c.PushBack(curr2.Value)
		curr2 = curr2.Next()
	}
	return c
}

// Split splits a list into two lists with approximately the same length,
// i.e. |a.Len() - b.Len()| <= 1.
func Split(l *list.List) (a, b *list.List) {
	a, b = new(list.List), new(list.List)
	if l.Len() < 2 {
		panic("length must be >= 2")
	}
	half := (l.Len() - 1) / 2
	curr := l.Front()
	for i := 0; uint(i) <= half; i++ {
		a.PushBack(curr.Value)
		curr = curr.Next()
	}
	for curr != nil {
		b.PushBack(curr.Value)
		curr = curr.Next()
	}
	return a, b
}

// MergeSort sorts a list using the merge-sort algorithm.
func MergeSort(l *list.List) *list.List {
	var c *list.List // sholdnt we split in len(l) lists or just 2
	if l.Len() <= 1 {
		return l
	}
	for i := l.Len(); i != 0; {
		i = i / 2
		c = Merge(Split(l))
	}
	return c
}

func main() {

}
