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
	l.PushBack(0)
	return l
}

func CaseNormal() *list.List {
	l := new(list.List)
	l.PushBack(0)
	l.PushBack(3)
	l.PushBack(1)
	l.PushBack(4)
	l.PushBack(90)
	return l
}
func CaseDuplicate() *list.List {
	l := new(list.List)
	l.PushBack(0)
	l.PushBack(3)
	l.PushBack(90)
	l.PushBack(3)
	l.PushBack(90)
	l.PushBack(0)
	return l
}
func CaseNegative() *list.List {
	l := new(list.List)
	l.PushBack(-5)
	l.PushBack(3)
	l.PushBack(1)
	l.PushBack(-98)
	l.PushBack(90)
	l.PushBack(102)
	return l
}

func TestMergeSort(t *testing.T) {
	MergeSort(CaseZero())
	MergeSort(CaseNormal())
	MergeSort(CaseNegative())
	MergeSort(CaseDuplicate())
	MergeSort(CaseNil())
}
