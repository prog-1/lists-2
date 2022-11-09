package main

import (
	"testing"

	"github.com/prog-1/list"
)

func TestMerge(t *testing.T) {
	for _, tc := range []struct {
		input1 []int
		input2 []int
		want   []int
	}{
		{[]int{1, 2}, []int{3, 4}, []int{1, 2, 3, 4}},
		{[]int{1, 3}, []int{3, 4}, []int{1, 3, 3, 4}},
		{[]int{1, 3}, []int{2, 4}, []int{1, 2, 3, 4}},
		{[]int{1, 4}, []int{2, 3}, []int{1, 2, 3, 4}},
		{[]int{1, 2, 3, 9}, []int{3, 4}, []int{1, 2, 3, 4, 9}},
	} {
		input1 := createListFromArray(tc.input1)
		input2 := createListFromArray(tc.input2)
		want := createListFromArray(tc.want)
		got := Merge(input1, input2)
		if !equal(want, *got) {
			t.Errorf("input %v, %v; want %v; got %v", tc.input1, tc.input2, tc.want, listToSlice(*got))
		}
	}
}

func TestSplit(t *testing.T) {
	for _, tc := range []struct {
		input []int
		want1 []int
		want2 []int
	}{
		{[]int{1, 2}, []int{1}, []int{2}},
		{[]int{1, 2, 3}, []int{1}, []int{2, 3}},
		{[]int{1, 3, 4, 5}, []int{1, 3}, []int{4, 5}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2}, []int{3, 4, 5}},
	} {
		input1 := createListFromArray(tc.input)
		want1 := createListFromArray(tc.want1)
		want2 := createListFromArray(tc.want2)

		got1, got2 := Split(&input1)
		if !equal(want1, *got1) || !equal(want2, *got2) {
			t.Errorf("input %v; want %v,%v ; got %v, %v", tc.input, tc.want1, tc.want2, listToSlice(*got1), listToSlice(*got2))
		}
	}
}

func TestMergeSort(t *testing.T) {
	for _, tc := range []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2}, []int{1, 2}},
		{[]int{1, 3, 2}, []int{1, 2, 3}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 5, 3, 4, 0}, []int{0, 1, 3, 4, 5}},
	} {
		input1 := createListFromArray(tc.input)
		want1 := createListFromArray(tc.want)

		got := MergeSort(&input1)
		if !equal(want1, *got) {
			t.Errorf("input %v; want %v ; got %v", tc.input, tc.want, listToSlice(*got))
		}
	}
}

func listToSlice(l list.List) []int {
	s := make([]int, l.Len())
	for i, n := 0, l.Front(); n != nil; n, i = n.Next(), i+1 {
		s[i] = n.Value
	}
	return s
}

func createListFromArray(s []int) (l list.List) {
	for i := range s {
		l.PushBack(s[i])
	}
	return l
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
