package main

import (
	"github.com/prog-1/list"
)

// Merge combines two sorted lists into a single sorted list.
func Merge(a, b *list.List) *list.List {
	l := new(list.List)
	eA, eB := a.Front(), b.Front()
	for eA != nil && eB != nil {
		if eA.Value < eB.Value {
			l.PushBack(eA.Value)
			eA = eA.Next()
		} else {
			l.PushBack(eB.Value)
			eB = eB.Next()
		}
	}
	for eA != nil {
		l.PushBack(eA.Value)
		eA = eA.Next()
	}
	for eB != nil {
		l.PushBack(eB.Value)
		eB = eB.Next()
	}
	return l
}

// Split splits a list into two lists with approximately the same length,
// i.e. |a.Len() - b.Len()| <= 1.
func Split(l *list.List) (a, b *list.List) {
	if l.Len() < 2 {
		panic("length must be >= 2")
	}
	a, b = new(list.List), new(list.List)
	e := l.Front()
	for i := uint(0); e != nil; i++ {
		if i < l.Len()/2 {
			a.PushBack(e.Value)
			e = e.Next()
		} else {
			b.PushBack(e.Value)
			e = e.Next()
		}
	}
	return a, b
}

// MergeSort sorts a list using the merge-sort algorithm.
func MergeSort(l *list.List) *list.List {
	if l.Len() == 1 {
		return l
	}
	a, b := Split(l)
	return Merge(MergeSort(a), MergeSort(b))
}
