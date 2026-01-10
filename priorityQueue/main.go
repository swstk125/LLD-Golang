package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	val1 int
	val2 int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].val2 < pq[j].val2
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func main() {
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{val1: 1, val2: 3})
	heap.Push(pq, &Item{val1: 2, val2: 1})
	heap.Push(pq, &Item{val1: 3, val2: 2})

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		fmt.Println(item.val1, item.val2)
	}
}
