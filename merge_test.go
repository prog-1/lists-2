package main

import (
	"testing"

	"github.com/prog-1/list"
)

func list1(e ...int) *list.List {
	l := new(list.List)
	for i := range e {
		l.PushBack(e[i])
	}
	return l
}

func TestMerge(t *testing.T) {
	for _, tc := range []struct {
		name string
		a, b *list.List
		want *list.List
	}{
		{"one of the lists is empty", list1(1, 2, 7, 8, 9), list1(), list1(1, 2, 7, 8, 9)},
		{"basic", list1(1, 2, 3), list1(4, 5, 6), list1(1, 2, 3, 4, 5, 6)},
		{"numbers repeat", list1(1, 3, 9, 9), list1(9, 11), list1(1, 2, 3, 9, 9, 9, 11)},
		{"merge", list1(1, 3, 9), list1(4, 5), list1(1, 3, 4, 5, 9)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			a, b, want := tc.a, tc.b, tc.want
			got := Merge(a, b)
			if !equal(got, want) {
				t.Errorf("got = %v, want = %v", got, want)
			}
		})
	}
}

func TestSplit(t *testing.T) {
	for _, tc := range []struct {
		name         string
		a            *list.List
		want1, want2 *list.List
	}{
		{"two numbers", list1(1, 1), list1(1), list1(1)},
		{"negative numbers", list1(9, 7, -3, 4, 0), list1(9, 7), list1(-3, 4, 0)},
		{"another case", list1(133, 1, 1, 3, 4), list1(133, 1), list1(1, 3, 4)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			a, want1, want2 := tc.a, tc.want1, tc.want2
			got1, got2 := Split(a)
			if !equal(got1, want1) {
				t.Errorf("got = %v, want = %v", got1, want1)
			} else if !equal(got2, want2) {
				t.Errorf("got = %v, want = %v", got2, want1)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	for _, tc := range []struct {
		name string
		a    *list.List
		want *list.List
	}{
		{"sorted", list1(1, 2, 3, 4, 6), list1(1, 2, 3, 4, 6)},
		{"unsorted", list1(1, 9, 6, 3, 2), list1(1, 2, 3, 6, 9)},
		{"unsorted, numbers repeat", list1(1, 9, 9, 6, 1, 6, 2, 1), list1(1, 1, 1, 2, 6, 9, 9)},
		{"negative numbers", list1(-4, 1, 5, -6), list1(-6, -4, 1, 5)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			a, want := tc.a, tc.want
			got := MergeSort(a)
			if !equal(got, want) {
				t.Errorf("got = %v, want = %v", got, want)
			}
		})
	}
}

func equal(a, b *list.List) bool {
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
