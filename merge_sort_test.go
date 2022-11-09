package main

import (
	"reflect"
	"testing"

	"github.com/prog-1/list"
)

func sliceToList(s []int) *list.List {
	l := &list.List{}
	for _, el := range s {
		l.PushBack(el)
	}
	return l
}
func listToSlice(l *list.List) []int {
	var s []int
	if l == nil {
		return nil
	}
	a := l.Front()
	for a != nil {
		s = append(s, a.Value)
		a = a.Next()
	}
	return s
}
func TestSplit(t *testing.T) {
	for _, tc := range []struct {
		a, want1, want2 []int
	}{

		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{1, 2, 3, 4}, []int{5, 6, 7, 8, 9}},
		{[]int{1, 2}, []int{1}, []int{2}},
		{[]int{1, 2, 54, 101, 123, 123}, []int{1, 2, 54}, []int{101, 123, 123}},
		{[]int{0, 1, 23, 54, 21, 21}, []int{0, 1, 23}, []int{54, 21, 21}},
	} {
		l := sliceToList(tc.a)
		a, b := Split(l)
		if got1, got2 := listToSlice(a), listToSlice(b); !reflect.DeepEqual(got1, tc.want1) || !reflect.DeepEqual(got2, tc.want2) {
			t.Errorf("Split(%v) = %v,%v, want %v,%v", tc.a, got1, got2, tc.want1, tc.want2)
		}
	}
}
