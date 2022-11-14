package main

import (
	"reflect"
	"testing"

	"github.com/prog-1/list"
)

// Merge combines two sorted lists into a single sorted list.
func Merge(a, b *list.List) *list.List {
	// Take first elements of each list
	// Compare what is smaller and save it to new result list
	// when one of the lists is empty, save the rest of the values in other list to result list and return it

	//edge cases
	if a.Front() == nil {
		return b
	} else if b.Front() == nil {
		return a
	}

	res := &list.List{}
	aEl, bEl := a.Front(), b.Front()
	for aEl != nil && bEl != nil {
		if aEl.Value < bEl.Value {
			res.PushBack(aEl.Value)
			aEl = aEl.Next()
		} else {
			res.PushBack(bEl.Value)
			bEl = bEl.Next()
		}
	}
	if aEl != nil {
		for aEl != nil {
			res.PushBack(aEl.Value)
			aEl = aEl.Next()
		}
	} else if bEl != nil {
		for bEl != nil {
			res.PushBack(bEl.Value)
			bEl = bEl.Next()
		}
	}
	return res
}

func TestMerge(t *testing.T) {
	for _, tc := range []struct {
		name string
		a    []int
		b    []int
		res  []int
	}{
		{"1", []int{1, 3, 5}, []int{2, 4, 6}, []int{1, 2, 3, 4, 5, 6}},
		{"2", []int{3, 6, 9}, []int{1, 2, 7, 10, 12}, []int{1, 2, 3, 6, 7, 9, 10, 12}},
		{"3", []int{1}, []int{2}, []int{1, 2}},
		{"4", []int{4}, []int{4}, []int{4, 4}},
		{"5", []int{1}, []int{}, []int{1}},
		{"5", []int{}, []int{}, []int{}},
	} {
		t.Run(tc.name, func(t *testing.T) {

			a := SliceToListBack(tc.a)
			b := SliceToListBack(tc.b)

			if !reflect.DeepEqual(ListToSlice(Merge(a, b)), tc.res) {
				t.Errorf("got = %v, want = %v", ListToSlice(Merge(a, b)), tc.res)
			}
		})
	}
}

//Converts slice into linked list using PushBack
func SliceToListBack(s []int) *list.List {
	l := &list.List{}
	for _, el := range s {
		l.PushBack(el)
	}
	return l
}

//Converts linked list into slice
func ListToSlice(l *list.List) []int {
	s := make([]int, 0)
	if l.Front() == nil {
		return s
	}
	s = append(s, l.Front().Value)
	el := l.Front().Next()
	for el != nil {
		s = append(s, el.Value)
		el = el.Next()
	}
	return s
}
