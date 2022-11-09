package main

import (
	"github.com/prog-1/list"

	"testing"
)

// Is it test - not at all

func CaseNil() *list.List {
	l := new(list.List)
	return l
}

func CaseZero() *list.List {
	l := new(list.List)
	for i := 0; i != 10000; i++ {
		l.PushBack(0)
	}
	return l
}

func CaseNormal() *list.List {
	l := new(list.List)
	for i := 0; i != 100000; i++ {
		l.PushBack(i)
	}
	return l
}
func CaseDuplicate() *list.List {
	l := new(list.List)
	for i := 0; i != 100000; i++ {
		l.PushBack(i)
		l.PushBack(i - 1)
		l.PushBack(i)
		l.PushBack(i + 1)
	}
	return l
}
func CaseNegative() *list.List {
	l := new(list.List)
	for i := -10000; i != 10000; i++ {
		l.PushBack(i)
		l.PushBack(i - 1)
		l.PushBack(i + 1)
	}
	return l
}

func TestMergeSort(t *testing.T) {
	MergeSort(CaseZero())
	MergeSort(CaseNormal())
	MergeSort(CaseNegative())
	MergeSort(CaseDuplicate())
	MergeSort(CaseNil())
}

// avg
func BenchmarkCaseZero(b *testing.B) { //2s
	MergeSort(CaseZero())
}
func BenchmarkN(b *testing.B) { //9s
	MergeSort(CaseNormal())
}
func BenchmarkNe(b *testing.B) {
	MergeSort(CaseNegative()) //3s
}
func BenchmarkNil(b *testing.B) { //1s
	MergeSort(CaseNil())
}
func BenchmarkD(b *testing.B) { // this benchmark should be slowest
	MergeSort(CaseDuplicate()) //2s avg???? in smaller numbers it was much slower...
}
