package main

import (
	"github.com/prog-1/list"
)

// Merge combines two sorted lists into a single sorted list.
func Merge(a, b *list.List) *list.List {
	curA, curB := a.Front(), b.Front()
	var list list.List
	var min *list.Element
	for {
		if curA != nil && curB != nil && curA.Value < curB.Value {
			min, curA = curA, curA.Next()
		} else if curA != nil && curB != nil && curA.Value > curB.Value {
			min, curB = curB, curB.Next()
		} else if curA != nil {
			min, curA = curA, curA.Next()
		} else if curB != nil {
			min, curB = curB, curB.Next()
		} else {
			break
		}
		list.PushBack(min.Value)
	}
	return &list
}

// Split splits a list into two lists with approximately the same length,
// i.e. |a.Len() - b.Len()| <= 1.
func Split(l *list.List) (a, b *list.List) {
	if l.Len() < 2 {
		panic("length must be >= 2")
	}
	// TODO
	return nil, nil
}

// MergeSort sorts a list using the merge-sort algorithm.
func MergeSort(l *list.List) *list.List {
	// TODO
	return nil
}

func main() {
}
