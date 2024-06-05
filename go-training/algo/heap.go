package algo

import (
	"container/heap"
)

// IntMinHeap defines a min-heap of integers.
type IntMinHeap []int

func (h IntMinHeap) Len() int           { return len(h) }
func (h IntMinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntMinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// IntMaxHeap defines a max-heap of integers.
type IntMaxHeap []int

func (h IntMaxHeap) Len() int           { return len(h) }
func (h IntMaxHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntMaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntMaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := &IntMinHeap{}
	heap.Init(h)

	for _, x := range nums {
		heap.Push(h, x)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	return heap.Pop(h).(int)
}

// https://leetcode.com/problems/top-k-frequent-elements/description/

type Element struct {
	value int
	freq  int
}

type ElementHeap []*Element

func (h ElementHeap) Len() int           { return len(h) }
func (h ElementHeap) Less(i, j int) bool { return h[i].freq < h[j].freq }
func (h ElementHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// push
func (h *ElementHeap) Push(e interface{}) {
	*h = append(*h, e.(*Element))
}

// pop
func (h *ElementHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// ElementHeap defines a min-heap of Element.

func topKFrequent(nums []int, k int) []int {
	h := &ElementHeap{}
	heap.Init(h)
	freq := make(map[int]int)

	for _, v := range nums {
		freq[v]++
	}

	for key, fr := range freq {
		heap.Push(h, &Element{
			value: key,
			freq:  fr,
		})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	rs := []int{}
	for h.Len() > 0 {
		e := heap.Pop(h).(*Element)
		rs = append(rs, e.value)
	}

	return rs
}
