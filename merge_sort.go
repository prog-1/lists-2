package main

import (
	"github.com/prog-1/list"
)

// Merge combines two sorted lists into a single sorted list.
func Merge(a, b *list.List) *list.List {
	l := new(list.List)
	a1 := a.Front()
	b1 := b.Front()
	for a1 != nil && b1 != nil {
		if a1.Value < b1.Value {
			l.PushBack(a1.Value)
			a1 = a1.Next()
		} else {
			l.PushBack(b1.Value)
			b1 = b1.Next()
		}
	}
	for a1 != nil {
		l.PushBack(a1.Value)
		a1 = a1.Next()
	}
	for b1 != nil {
		l.PushBack(b1.Value)
		b1 = b1.Next()
	}
	return l
}

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
		a.PushBack(l1.Value)
		l1 = l1.Next()
	}
	b = new(list.List)
	for ; i <= n && l1.Next() != nil; i++ {
		b.PushBack(l1.Value)
		l1 = l1.Next()
	}
	b.PushBack(l1.Value)
	return a, b
}

// MergeSort sorts a list using the merge-sort algorithm.
func MergeSort(l *list.List) *list.List {
	if l.Len() <= 2 {
		if l.Len() == 1 {
			list.Print(l)
			return l
		}
		a, b := Split(l)
		return Merge(a, b)
	}
	a, b := Split(l)
	l = Merge(MergeSort(a), b)
	a, b = Split(l)
	l = Merge(MergeSort(a), b)
	return l
}
func stl(s []int) *list.List {
	l := &list.List{}
	for _, el := range s {
		l.PushBack(el)
	}
	return l
}
func main() {
	l := MergeSort(stl([]int{2, 1}))
	list.Print(l)
}
