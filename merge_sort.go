package main

import (
	"github.com/prog-1/list"
)

// Merge combines two sorted lists into a single sorted list.
func Merge(a, b *list.List) *list.List {
	// TODO
	return nil
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
