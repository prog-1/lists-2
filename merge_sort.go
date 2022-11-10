package main

import (
	"github.com/prog-1/list"
)

type T int

type Element struct {
	Value T
	next  *Element
}

func NewElement(v T) *Element {
	return &Element{Value: v}
}

func (e *Element) Next() *Element {
	return e.next
}

type List struct {
	head, tail *Element
	len        uint
}

func (l *List) Front() *Element {
	return l.head
}

func (l *List) Back() *Element {
	return l.tail
}

// Merge combines two sorted lists into a single sorted list.
func Merge(a, b *list.List) *list.List {
	a1 := a.Front()
	b1 := b.Front()
	if a1 == nil {
		return b
	}
	if b1 == nil {
		return b
	}
	if a1.Value < b1.Value {
		a1.Next() = Merge(a1.Next(), b1.Value)
		return a
	}
	b1.Next() = Merge(a1, b1.Next())
	return b
}

func (l *List) PushBack(v T) {
	e := NewElement(v)
	if l.len == 0 {
		l.head, l.tail = e, e
	} else {
		l.tail.next = e
		l.tail = e
	}
	l.len++
}

// Split splits a list into two lists with approximately the same length,
// i.e. |a.Len() - b.Len()| <= 1.
func Split(l *list.List) (a, b *list.List) {
	if l.Len() < 2 {
		panic("length must be >= 2")
	}
	a = new(list.List)
	b = new(list.List)
	var xd uint
	for e := l.Front(); e != nil; e = e.Next() {
		xd++
		if xd < l.Len()/2 {
			a.PushBack(e.Value)
		} else {
			b.PushBack(e.Value)
		}
	}
	return a, b
}

// MergeSort sorts a list using the merge-sort algorithm.
func MergeSort(l *list.List) *list.List {
	// TODO
	return nil
}

func main() {

}
