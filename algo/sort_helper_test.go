package algo

import (
	"testing"
)

func TestSortHelper_MergeSort(t *testing.T) {
	sh := SortHelper{}
	data := []int{8, 3, 1, 9, 2}
	sh.MergeSort(data, 0, len(data))
	t.Log(data)
}

func TestSortHelper_MergeReOrder(t *testing.T) {
	var count int
	sh := SortHelper{}
	//data := []int{7,5,6,4}
	data := []int{5,1,4,2,8}
	reorderList := make([][]int, 0)
	sh.MergeReOrder(data, 0, len(data), &count, &reorderList)
	t.Log("count", count)
	t.Log("reorderList", reorderList)
}

func TestSortHelper_MergeVSort(t *testing.T) {
	sh := SortHelper{}
	data := []int{8, 3, 1, 9, 2}
	res := sh.MergeVSort(data)
	t.Log(res)
}

func TestSortHelper_QuickSort(t *testing.T) {
	sh := SortHelper{}
	data := []int{8, 3, 1, 9, 2, 5}
	sh.QuickSort(data)
	t.Log(data)
}

func TestSortHelper_TopQuickSort(t *testing.T) {
	sh := SortHelper{}
	data := []int{8, 3, 1, 9, 2, 5}
	desc := sh.TopQuickSort(data, 3, true)
	asc := sh.TopQuickSort(data, 2, false)
	t.Log(data)
	t.Log(desc)
	t.Log(asc)
}

func TestBuildTopKFrequentHeap(t *testing.T) {
	BuildTopKFrequentHeap(2)
}
