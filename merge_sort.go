package main

import (
	"fmt"

	"github.com/prog-1/list"
)

// Merge combines two sorted lists into a single sorted list.
// func Merge(a, b *list.List) (c *list.List) {
// 	p := &c
// 	// if a == nil {
// 	// 	return b
// 	// } else if b == nil {
// 	// 	return a
// 	// }
// 	// n := b.Len() + a.Len()
// 	// switch {
// 	// case a == nil:
// 	// 	return b
// 	// case b == nil:
// 	// 	return a
// 	// case a1.Value < b1.Value:
// 	// 	a = Merge(a.Next(), b)
// 	// 	return a
// 	// default:
// 	// 	b = Merge(a, b1.Next())
// 	// 	return b
// 	// }
// 	a1 := a.Front()
// 	b1 := b.Front()
// 	for a != nil && b != nil {
// 		if a1.Value < b1.Value {
// 			*p, a = a, a1.
// 		} else {
// 			p.Next = l2
// 			l2 = l2.Next
// 		}
// 		return nil
// 	}
// }

// Split splits a list into two lists with approximately the same length,
// i.e. |a.Len() - b.Len()| <= 1.
func Split(l *list.List) (a, b *list.List) {
	if l.Len() < 2 {
		panic("length must be >= 2")
	}
	a = new(list.List)
	n := int(l.Len())
	i := 0
	l1 := l.Front()
	for ; i < n/2 && l1.Next() != nil; i++ {
		fmt.Println(l1.Value)
		a.PushBack(l1.Value)
		l1 = l1.Next()
	}
	b = new(list.List)
	for ; i <= n && l1.Next() != nil; i++ {
		fmt.Println(l1.Value)
		b.PushBack(l1.Value)
		l1 = l1.Next()
	}
	b.PushBack(l1.Value)
	return a, b
}

// MergeSort sorts a list using the merge-sort algorithm.
func MergeSort(l *list.List) *list.List {
	// TODO
	return nil
}

func main() {
}
