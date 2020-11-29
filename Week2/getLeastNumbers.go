package main

import (
	"container/heap"
	"sort"
)

// O(nlogn) 44ms
func getLeastNumbers(arr []int, k int) []int {
	if k > len(arr) {
		return nil
	}
	sort.Ints(arr)
	return arr[:k]
}

// O(nlogn) :维护了一个大小为n的小顶堆 44ms
func getLeastNumbersMinHeap(arr []int, k int) []int {
	h := &IntHeap{}
	for _, info := range arr {
		heap.Push(h, info)
	}
	res := make([]int, 0, len(arr))
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(h).(int))
	}
	return res
}

// O(nlogk) :维护了一个大小为k的大顶堆 36ms
func getLeastNumbersMaxHeap(arr []int, k int) []int {
	if len(arr) == 0 || k == 0 {
		return nil
	}
	h := &IntHeap{}
	for i := 0; i < k; i++ {
		heap.Push(h, arr[i])
	}
	for i := k; i < len(arr); i++ {
		// 维持堆大小为k
		if arr[i] < h.peer() {
			heap.Pop(h)
			heap.Push(h, arr[i])
		}
	}
	// res = make([]int, k)
	return *h
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// O(n) 限制条件 28ms
// 0 <= k <= arr.length <= 10000
// 0 <= arr[i] <= 10000
// 不论多少数据进入筛选都只循环10000次，适用于数据范围有限制，数据量较大的数据筛选
func getLeastNumbersLimit(arr []int, k int) []int {
	if k == 0 || len(arr) == 0 {
		return []int{}
	}
	// 统计每个数字出现的次数
	var counter [10001]int
	for _, num := range arr {
		counter[num]++
	}
	// 根据counter数组从头找出k个数作为返回结果
	res := make([]int, k)
	idx := 0
	for num := 0; num < len(counter); num++ {
		for counter[num] > 0 && idx < k {
			res[idx] = num
			idx++
			counter[num]--
		}
		if idx == k {
			break
		}
	}
	return res
}

// O(n) quick-sort mind
func getLeastNumbers(arr []int, k int) []int {
	if k >= len(arr) {
		return arr
	}
	return quickselect(arr, 0, len(arr)-1, k)
}

func quickselect(arr []int, left, right int, k int) []int {
	if left < right {
		index := partition(arr, left, right)
		if index == k {
			return arr[:k]
		} else if index > k {
			return quickselect(arr, left, index-1, k)
		} else if index < k {
			return quickselect(arr, index+1, right, k)
		}
	}
	return arr[:k]
}

func partition(num []int, l, r int) int {
	poivt := num[l]
	index := l + 1
	for i := index; i <= r; i++ {
		if num[i] < poivt {
			num[i], num[index] = num[index], num[i]
			index++
		}
	}
	num[l], num[index-1] = num[index-1], num[l]
	return index - 1
}
