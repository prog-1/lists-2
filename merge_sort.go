package main

import (
	"fmt"

	"github.com/prog-1/list"
)

// Merge combines two sorted lists into a single sorted list.
func Merge(a, b *list.List) *list.List {
	var l list.List
	curA, curB := a.Front(), b.Front()
	var min *list.Element
	for {
		if curA != nil && curB != nil && curA.Value < curB.Value {
			min = curA
			curA = curA.Next()
		} else if curB != nil {
			min = curB
			curB = curB.Next()
		} else if curA != nil {
			min = curA
			curA = curA.Next()
		} else {
			break
		}
		l.PushBack(min.Value)
	}

	return &l
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
	var a list.List
	a.PushBack(1)
	a.PushBack(3)
	var b list.List
	b.PushBack(2)
	b.PushBack(4)
	l := Merge(&a, &b)
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
