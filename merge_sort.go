package main

import (
	"github.com/prog-1/list"
)

// Merge combines two sorted lists into a single sorted list.

// tried to write recursive merge
func Merge2(a, b *list.Element) (*list.Element, *list.List) {
	var res *list.List
	if a == nil {
		res.PushFront(b.Value)
		return b, res
	}
	if b == nil {
		res.PushFront(a.Value)
		return a, res
	}
	if a.Value < b.Value {
		res.PushFront(a.Value)
		a = a.Next()
		a, res = Merge2(a, b)
		return a, res
	}
	res.PushFront(b.Value)
	b = b.Next()
	b, res = Merge2(a, b)
	return b, res
}

func Merge(a, b *list.List) *list.List {
	a1 := a.Front()
	b1 := b.Front()
	_, res := Merge2(a1, b1)
	return res
}

// Split splits a list into two lists with approximately the same length,
// i.e. |a.Len() - b.Len()| <= 1.
func Split(l *list.List) (a, b *list.List) {
	if l.Len() < 2 {
		return
	}
	res1, res2 := new(list.List), new(list.List)
	el := l.Front()
	p := l.Len() / 2
	var i uint
	for ; el != nil && i < l.Len(); i++ {
		if i < p {
			res1.PushBack(el.Value)
			el = el.Next()
		} else {
			res2.PushBack(el.Value)
			el = el.Next()
		}
	}
	return res1, res2
}

// MergeSort sorts a list using the merge-sort algorithm.
func MergeSort(l *list.List) *list.List {
	if l.Len() == 1 {
		return l
	}
	a, b := Split(l)
	return Merge(MergeSort(a), MergeSort(b))
}

func main() {
}
