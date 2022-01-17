package algo

import (
	"container/heap"
	"fmt"
	"math/rand"
	"time"
)

type SortHelper struct {
}

func (sh SortHelper) MergeVSort(data []int) []int {
	if len(data) < 2 {
		return data
	}
	mid := len(data) >> 1
	left := sh.MergeVSort(data[0:mid])
	fmt.Println("mid", mid, left, data)
	right := sh.MergeVSort(data[mid:])
	fmt.Println(left, right)
	return bing(left, right)
}

func bing(left, right []int) []int {
	var res []int
	j := 0
	k := 0
	for j < len(left) && k < len(right) {
		if left[j] > right[k] {
			res = append(res, right[k])
			k++
		} else {
			res = append(res, left[j])
			j++
		}
	}
	res = append(res, left[j:]...)
	res = append(res, right[k:]...)
	return res
}

func (sh SortHelper) MergeSort(data []int, start, end int) {
	if end-start > 1 {
		mid := start + (end-start+1)>>1
		sh.MergeSort(data, start, mid)
		sh.MergeSort(data, mid, end)
		fmt.Println(start, mid, end)
		merge(data, start, mid, end)
	}
}

func (sh SortHelper) QuickSort(data []int) {
	quickSort(data, 0, len(data)-1)
}

func quickSort(data []int, l, r int) {
	if l < r {
		mid := quickSortPartition(data, l, r)
		fmt.Println("mid: ", mid, l, r)
		quickSort(data, l, mid-1)
		quickSort(data, mid+1, r)
	}
}

func (sh SortHelper) TopQuickSort(data []int, top int, isDesc bool) []int {
	if top > len(data) {
		return nil
	}
	// 最大Top k
	if isDesc {
		topQuickSort(data, 0, len(data)-1, len(data)-top)
		return data[(len(data) - top):]
	}
	// 最小Top k
	topQuickSort(data, 0, len(data)-1, top-1)
	return data[:top]
}

func topQuickSort(data []int, l, r, top int) {
	if l < r {
		mid := quickSortPartition(data, l, r)
		fmt.Println("mid: ", mid, l, r, data)
		if top > mid {
			topQuickSort(data, mid+1, r, top)
		} else if top < mid {
			topQuickSort(data, l, mid-1, top)
		} else if top == mid {
			return
		}
	}
}

func quickSortPartition(data []int, l, r int) int {
	target := data[l]
	for l < r {
		for l < r && data[r] >= target {
			r--
		}
		data[l], data[r] = data[r], data[l]
		for l < r && data[l] <= target {
			l++
		}
		data[l], data[r] = data[r], data[l]
	}
	return l
}

func merge(data []int, start, mid, end int) {
	lSize := mid - start
	rSize := end - mid
	s := end - start
	temp := make([]int, 0, s)
	l, r := 0, 0
	for l < lSize && r < rSize {
		lValue := data[start+l]
		rValue := data[mid+r]
		if lValue < rValue {
			temp = append(temp, lValue)
			l++
		} else {
			temp = append(temp, rValue)
			r++
		}
	}
	if l < lSize {
		temp = append(temp, data[start+l:mid]...)
	}
	if r < rSize {
		temp = append(temp, data[mid+r:end]...)
	}

	for i := 0; i < s; i++ {
		data[start+i] = temp[i]
	}
	fmt.Println("temp: ", temp, data)
}

type HeapObj [][2]int

func (ho HeapObj) Len() int {
	return len(ho)
}

func (ho HeapObj) Less(i, j int) bool {
	// 建立最小堆
	return ho[i][1] < ho[j][1]
}

func (ho HeapObj) Swap(i, j int) {
	ho[i], ho[j] = ho[j], ho[i]
}

func (ho *HeapObj) Push(x interface{}) {
	*ho = append(*ho, x.([2]int))
}

func (ho *HeapObj) Pop() interface{} {
	last := (*ho)[ho.Len()-1]
	*ho = (*ho)[:ho.Len()-1]
	return last
}

// https://leetcode-cn.com/problems/top-k-frequent-elements/
func BuildTopKFrequentHeap(maxLen int) {
	rand.Seed(time.Now().Unix())
	obj := &HeapObj{}
	heap.Init(obj)
	for i := 0; i < 5; i++ {
		heap.Push(obj, [2]int{i + 1, rand.Intn(20)})
		// 识别超出，进行出栈处理
		if obj.Len() > maxLen {
			res := heap.Pop(obj)
			data := res.([2]int)
			fmt.Printf("pop less data: %d:%d \n", data[0], data[1])
		}
	}
	fmt.Println("rest data: ", *obj)
	for obj.Len() > maxLen {
		res := heap.Pop(obj)
		data := res.([2]int)
		fmt.Printf("pop less data: %d:%d \n", data[0], data[1])
	}
	maxRes := make([]int, maxLen)
	for i := 0; i < maxLen; i++ {
		res := heap.Pop(obj)
		data := res.([2]int)
		maxRes[maxLen-1-i] = data[0]
	}
	fmt.Println(maxRes)
}

type HeapPairObj struct {
	data     [][2]int
	numList  []int
	numList2 []int
}

func (ho HeapPairObj) Len() int {
	return len(ho.data)
}

func (ho HeapPairObj) Less(i, j int) bool {
	// 建立最小堆
	first, second := ho.data[i], ho.data[j]
	return ho.numList[first[0]]+ho.numList2[first[1]] < ho.numList[second[0]]+ho.numList2[second[1]]
}

func (ho HeapPairObj) Swap(i, j int) {
	ho.data[i], ho.data[j] = ho.data[j], ho.data[i]
}

func (ho *HeapPairObj) Push(x interface{}) {
	ho.data = append(ho.data, x.([2]int))
}

func (ho *HeapPairObj) Pop() interface{} {
	last := ho.data[ho.Len()-1]
	ho.data = ho.data[:ho.Len()-1]
	return last
}

func kSmallestPairs(nums1, nums2 []int, k int) [][]int {
	res := make([][]int, 0)
	n1 := len(nums1)
	n2 := len(nums2)
	obj := &HeapPairObj{
		data:     make([][2]int, 0, k),
		numList:  nums1,
		numList2: nums2,
	}
	heap.Init(obj)
	for i := 0; i < k && i < n1; i++ {
		heap.Push(obj, [2]int{i, 0})
	}
	for obj.Len() > 0 && len(res) < k {
		target := heap.Pop(obj)
		data := target.([2]int)
		res = append(res, []int{obj.numList[data[0]], obj.numList2[data[1]]})
		// 下标加一不会超出边界
		if data[1]+1 < n2 {
			// 补进到最小堆里, 检测是否有最小值组合
			heap.Push(obj, [2]int{data[0], data[1] + 1})
		}
	}
	return res
}
