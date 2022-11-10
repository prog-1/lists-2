package main

import (
	"testing"

	"github.com/prog-1/list"
)

func initList(e ...int) *list.List {
	l := new(list.List)
	for i := range e {
		l.PushBack(e[i])
	}
	return l
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

func TestMerge(t *testing.T) {
	for _, tc := range []struct {
		initAF, initBF, wantF *list.List
	}{
		{initList(), initList(), initList()},
		{initList(1), initList(3), initList(1, 3)},
		{initList(-2, -3, 0), initList(1, 11, -20), initList(-20, -3, -2, 0, 1, 11)},
		{initList(1, 2), initList(3, 4), initList(1, 2, 3, 4)},
		{initList(1, 4), initList(2, 3, 5), initList(1, 2, 3, 4, 5)},
		{initList(1, 3), initList(2, 3, 4), initList(1, 2, 3, 3, 4)},
	} {
		a, b, want := tc.initAF, tc.initBF, tc.wantF
		got := Merge(a, b)
		if !equal(got, want) {
			t.Errorf("got = %v, want = %v", got, want)
		}
	}
}

func TestSplit(t *testing.T) {
	for _, tc := range []struct {
		initF, wantAF, wantBF *list.List
	}{
		{initList(1, 1), initList(1), initList(1)},
		{initList(2, -3, 0, 99), initList(2, -3), initList(0, 99)},
		{initList(2, 1, 3, 4), initList(2, 1), initList(3, 4)},
		{initList(2, 1, 3, 4, 0), initList(2, 1), initList(3, 4, 0)},
		{initList(2, 1, 1, 3, 4), initList(2, 1), initList(1, 3, 4)},
	} {
		l, wantA, wantB := tc.initF, tc.wantAF, tc.wantBF
		gotA, gotB := Split(l)
		if !equal(gotA, wantA) {
			t.Errorf("gotA = %v, wantA = %v", gotA, wantA)
		} else if !equal(gotB, wantB) {
			t.Errorf("gotB = %v, wantB = %v", gotB, wantB)
		}

	}
}

func TestMergeSort(t *testing.T) {
	for _, tc := range []struct {
		initF, wantF *list.List
	}{
		{initList(), initList()},
		{initList(1), initList(1)},
		{initList(1000, -7), initList(-7, 1000)},
		{initList(2, 1, -1), initList(-1, 1, 2)},
		{initList(1, 2, 0, 5, 4, 3), initList(0, 1, 2, 3, 4, 5)},
		{initList(-1, 2, 0, 5, 4, 3, 6), initList(-1, 0, 2, 3, 4, 5, 6)},
	} {
		got, want := tc.initF, tc.wantF
		got = MergeSort(got)
		if !equal(got, want) {
			t.Errorf("got = %v, want = %v", got, want)
		}
	}
}
