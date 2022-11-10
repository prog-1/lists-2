package main

import (
	"github.com/prog-1/list"
)

// Merge combines two sorted lists into a single sorted list.
func Merge(a, b *list.List) *list.List {
	h := new(list.List)
	fa, fb := a.Front(), b.Front()
	for fa != nil || fb != nil {
		if fb.Value > fa.Value {
			h.PushBack(fa.Value)
			fa = fa.Next()
		} else {
			h.PushBack(fb.Value)
			fb = fb.Next()
		}
	}
	return h
}

// Split splits a list into two lists with approximately the same length,
// i.e. |a.Len() - b.Len()| <= 1.0
func Split(l *list.List) (a, b *list.List) {
	if l.Len() < 2 {
		return
	}
	out1, out2 := new(list.List), new(list.List)
	c := l.Front()
	for cnt := l.Len() / 2; cnt > 0; cnt-- {
		out1.PushBack(c.Value)
		c = c.Next()
	}
	for c != nil {
		out2.PushBack(c.Value)
		c = c.Next()
	}
	return out1, out2
}

// MergeSort sorts a list using the merge-sort algorithm.
func MergeSort(l *list.List) *list.List {
	if l.Len() == 1 {
		return l
	}
	a, b := Split(l)
	return Merge(MergeSort(a), MergeSort(b))
}
