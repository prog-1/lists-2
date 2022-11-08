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

	// Getting middle element:
	var midEl *list.Element = l.Front()
	mid := (int)(l.Len() / 2)
	for i := 0; i < mid; i++ {
		midEl = midEl.Next()
	}

	// Populating a:
	a = &list.List{}
	elA := l.Front()
	for i := 0; i < mid; i, elA = i+1, elA.Next() {
		a.PushBack(elA.Value)
	}

	// Populating b:
	b = &list.List{}
	for elB := elA; elB != nil; elB = elB.Next() {
		b.PushBack(elB.Value)
	}

	return
}

// MergeSort sorts a list using the merge-sort algorithm.
func MergeSort(l *list.List) *list.List {
	if l.Len() == 1 {
		return l
	}
	a, b := Split(l)
	a = MergeSort(a)
	b = MergeSort(b)
	return Merge(a, b)
}

func main() {
	var l list.List
	l.PushBack(1)
	l.PushBack(3)
	l.PushBack(2)
	l.PushBack(4)
	l.PushBack(2)
	l.PushBack(4)
	//var b list.List
	//b.PushBack(2)
	//b.PushBack(4)
	//l := Merge(&a, &b)
	// a, b := Split(&l)
	// for i := a.Front(); i != nil; i = i.Next() {
	// 	fmt.Println(i.Value)
	// }
	// fmt.Println("-----------------")
	// for i := b.Front(); i != nil; i = i.Next() {
	// 	fmt.Println(i.Value)
	// }
	r := MergeSort(&l)
	for i := r.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
