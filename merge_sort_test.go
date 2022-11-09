package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/prog-1/list"
)

func List(e ...int) func() list.List {
	return func() (l list.List) {
		for i := range e {
			l.PushBack(e[i])
		}
		return l
	}
}

func equal(a, b list.List) bool {
	if a.Len() != b.Len() {
		return false
	}
	for a, b := a.Front(), b.Front(); a != nil; a, b = a.Next(), b.Next() {
		if a.Value != b.Value {
			return false
		}
	}
	return true
}

func format(l list.List) string {
	var b bytes.Buffer
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Fprintf(&b, "[%v]->", e.Value)
	}
	b.WriteString("|")
	return b.String()
}

func TestMerge(t *testing.T) {
	for _, tc := range []struct {
		name       string
		a, b, want func() list.List
	}{
		{"1", List(1), List(1), List(1, 1)},
		{"2", List(1), List(2), List(1, 2)},
		{"3", List(1, 1), List(1), List(1, 1, 1)},
		{"4", List(1), List(1, 1), List(1, 1, 1)},
		{"5", List(1, 1), List(1, 1), List(1, 1, 1, 1)},
		{"6", List(1, 2), List(1, 2), List(1, 1, 2, 2)},
		{"7", List(1, 2), List(3, 4), List(1, 2, 3, 4)},
		{"8", List(1, 2, 4), List(1, 3, 4), List(1, 1, 2, 3, 4, 4)},
		{"9", List(1, 2, 3, 123), List(4, 5, 6, 51), List(1, 2, 3, 4, 5, 6, 51, 123)},
		{"10", List(1, 2, 3, 7, 123), List(4, 5, 6, 51), List(1, 2, 3, 4, 5, 6, 7, 51, 123)},
	} {
		t.Run(fmt.Sprint(tc.name), func(t *testing.T) {
			l1, l2 := tc.a(), tc.b()
			got := Merge(&l1, &l2)
			if !equal(*got, tc.want()) {
				t.Errorf("Merge(%v %v) = %v, want %v", format(tc.a()), format(tc.b()), format(*got), format(tc.want()))
			}
		})
	}
}

func TestSplit(t *testing.T) {
	for _, tc := range []struct {
		name                string
		input, want1, want2 func() list.List
	}{
		{"1", List(1, 1), List(1), List(1)},
		{"2", List(1, 2), List(1), List(2)},
		{"3", List(1, 2, 3), List(1, 2), List(3)},
		{"4", List(1, 2, 4), List(1, 2), List(4)},
		{"5", List(3, 1, 2), List(3, 1), List(2)},
		{"6", List(1, 2, 3, 4), List(1, 2), List(3, 4)},
		{"7", List(5, 12, 4, 27), List(5, 12), List(4, 27)},
		{"8", List(6, 23, 14, 127, 15), List(6, 23, 14), List(127, 15)},
		{"9", List(74, 100, 24, 219, 199, 515), List(74, 100, 24), List(219, 199, 515)},
	} {
		t.Run(fmt.Sprint(tc.name), func(t *testing.T) {
			l := tc.input()
			got1, got2 := Split(&l)
			if !equal(*got1, tc.want1()) || !equal(*got2, tc.want2()) {
				t.Errorf("Split(%v) = %v %v, want %v %v", format(tc.input()), format(*got1), format(*got2), format(tc.want1()), format(tc.want2()))
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	for _, tc := range []struct {
		name        string
		input, want func() list.List
	}{
		{"1", List(0), List(0)},
		{"2", List(1), List(1)},
		{"3", List(2), List(2)},
		{"4", List(1, 2), List(1, 2)},
		{"5", List(2, 1), List(1, 2)},
		{"6", List(3, 1), List(1, 3)},
		{"7", List(1, 2, 3), List(1, 2, 3)},
		{"8", List(1, 3, 2), List(1, 2, 3)},
		{"9", List(2, 1, 3), List(1, 2, 3)},
		{"10", List(2, 3, 1), List(1, 2, 3)},
		{"11", List(3, 2, 1), List(1, 2, 3)},
		{"12", List(3, 1, 2), List(1, 2, 3)},
		{"13", List(1, 2, 3, 5), List(1, 2, 3, 5)},
		{"14", List(1, 2, 5, 3), List(1, 2, 3, 5)},
		{"15", List(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12), List(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12)},
		{"16", List(12, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0), List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12)},
		{"17", List(13, 10, 9, 3, 7, 6, 5, 4, 8, 2, 1, 0), List(0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 13)},
		{"18", List(8, 7439, 4, 8, 1, 7, 9, 90, 13, 3, 123, 51, 8, 9), List(1, 3, 4, 7, 8, 8, 8, 9, 9, 13, 51, 90, 123, 7439)},
	} {
		t.Run(fmt.Sprint(tc.name), func(t *testing.T) {
			l := tc.input()
			got := MergeSort(&l)
			if !equal(*got, tc.want()) {
				t.Errorf("MergeSort(%v) = %v, want %v", format(tc.input()), format(*got), format(tc.want()))
			}
		})
	}
}
