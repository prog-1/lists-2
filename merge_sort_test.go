package main

import (
	"bytes"
	"fmt"
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

func format(l *list.List) string {
	var b bytes.Buffer
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Fprintf(&b, "->[%v]", e.Value)
	}
	b.WriteString("->|")
	return b.String()
}

func TestMerge(t *testing.T) {
	for _, tc := range []struct {
		name                  string
		initAF, initBF, wantF *list.List
	}{
		{"case-1", initList(1, 2), initList(3, 4), initList(1, 2, 3, 4)},
		{"case-2", initList(1, 4), initList(2, 3, 5), initList(1, 2, 3, 4, 5)},
		{"case-3", initList(1, 3), initList(2, 3, 4), initList(1, 2, 3, 3, 4)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			a, b, want := tc.initAF, tc.initBF, tc.wantF
			got := Merge(a, b)
			if !equal(got, want) {
				t.Errorf("got = %v, want = %v", format(got), format(want))
			}
		})
	}
}

func TestSplit(t *testing.T) {
	for _, tc := range []struct {
		name                  string
		initF, wantAF, wantBF *list.List
	}{
		{"case-1", initList(2, 1, 3, 4), initList(2, 1), initList(3, 4)},
		{"case-2", initList(2, 1, 3, 4, 0), initList(2, 1), initList(3, 4, 0)},
		{"case-3", initList(2, 1, 1, 3, 4), initList(2, 1), initList(1, 3, 4)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			l, wantA, wantB := tc.initF, tc.wantAF, tc.wantBF
			gotA, gotB := Split(l)
			if !equal(gotA, wantA) {
				t.Errorf("gotA = %v, wantA = %v", format(gotA), format(wantA))
			} else if !equal(gotB, wantB) {
				t.Errorf("gotB = %v, wantB = %v", format(gotB), format(wantB))
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	for _, tc := range []struct {
		name         string
		initF, wantF *list.List
	}{
		{"case-1", initList(2, 1, -1), initList(-1, 1, 2)},
		{"case-2", initList(1, 2, 0, 5, 4, 3), initList(0, 1, 2, 3, 4, 5)},
		{"case-3", initList(-1, 2, 0, 5, 4, 3, 6), initList(-1, 0, 2, 3, 4, 5, 6)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, want := tc.initF, tc.wantF
			got = MergeSort(got)
			if !equal(got, want) {
				t.Errorf("got = %v, want = %v", format(got), format(want))
			}
		})
	}
}
