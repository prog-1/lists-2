package main

import (
	"fmt"

	"github.com/prog-1/list"
)

// Merge combines two sorted lists into a single sorted list.
func Merge(a, b *list.List) *list.List {
	var l list.List
	n, k := a.Front(), b.Front()
	for n != nil || k != nil {
		if k == nil || n != nil && n.Value <= k.Value {
			l.PushBack(n.Value)
			n = n.Next()
		} else if n == nil || k != nil && k.Value <= n.Value {
			l.PushBack(k.Value)
			k = k.Next()
		}
	}
	return &l
}

// Split splits a list into two lists with approximately the same length,
// i.e. |a.Len() - b.Len()| <= 1.
func Split(l *list.List) (a, b *list.List) {
	if l.Len() < 2 {
		panic("length must be >= 2")
	}
	a, b = new(list.List), new(list.List)
	for n := l.Front(); n != nil; n = n.Next() {
		if a.Len() <= (l.Len()-1)/2 {
			a.PushBack(n.Value)
		} else {
			b.PushBack(n.Value)
		}
	}
	return
}

// MergeSort sorts a list using the merge-sort algorithm.
func MergeSort(l *list.List) *list.List {
	if l.Len() < 2 {
		return l
	}
	s1, s2 := Split(l)
	s1 = MergeSort(s1)
	s2 = MergeSort(s2)
	return Merge(s1, s2)
}

func main() {
	var a, b list.List
	for i := 1; i < 4; i++ {
		a.PushBack(i)
		b.PushBack(i + 3)
	}
	a.PushBack(123)
	b.PushBack(51)
	list.Print(&a)
	list.Print(&b)
	res := Merge(&a, &b)
	list.Print(res)

	fmt.Println()
	var l list.List
	s := []int{4, 8, 1, 7, 9, 13, 3}
	for _, i := range s {
		l.PushBack(i)
	}
	l.PushFront(8)
	list.Print(&l)
	n, k := Split(&l)
	list.Print(n)
	list.Print(k)

	fmt.Println()
	res = MergeSort(&l)
	list.Print(res)
}
