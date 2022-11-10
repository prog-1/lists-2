package main

import (
	"fmt"

	"github.com/prog-1/list"
)

// Merge combines two sorted lists into a single sorted list.
func Merge(a, b list.List) *list.List {
	c := &list.List{}
	elementA := a.Front()
	elementB := b.Front()
	for elementA != nil && elementB != nil {
		if elementA.Value > elementB.Value {
			c.PushBack(elementB.Value)
			elementB = elementB.Next()
		} else {
			c.PushBack(elementA.Value)
			elementA = elementA.Next()
		}
	}
	if elementA == nil {
		for ; elementB != nil; elementB = elementB.Next() {
			c.PushBack(elementB.Value)
		}
	} else {
		for ; elementA != nil; elementA = elementA.Next() {
			c.PushBack(elementA.Value)
		}
	}
	return c
}

// Split splits a list into two lists with approximately the same length,
// i.e. |a.Len() - b.Len()| <= 1.
func Split(l *list.List) (*list.List, *list.List) {
	if l.Len() < 2 {
		panic("length must be >= 2")
	}
	tmp := l.Front()
	a := &list.List{}
	b := &list.List{}

	for i := l.Len() / 2; i > 0; i-- {
		a.PushBack(tmp.Value)
		tmp = tmp.Next()
	}
	for tmp != nil {
		b.PushBack(tmp.Value)
		tmp = tmp.Next()
	}
	return a, b
}

// MergeSort sorts a list using the merge-sort algorithm.
func MergeSort(l *list.List) *list.List {
	if l.Len() == 1 {
		return l
	}
	a, b := Split(l)
	return Merge(*MergeSort(a), *MergeSort(b))
}

// func printlist(l ...*list.List) {
// 	for i, v := range l {
// 		for n := v.Front(); n != nil; n = n.Next() {
// 			fmt.Println(i, n)
// 		}
// 	}
// }

func main() {
	var a list.List
	a.PushBack(2)
	a.PushBack(5)
	a.PushBack(4)
	a.PushBack(1)
	a.PushBack(8)

	l := MergeSort(&a)
	for n := l.Front(); n != nil; n = n.Next() {
		fmt.Println(n)
	}
}
