package main

import (
	"container/heap"
	"container/list"
)
import "fmt"

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

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

func main() {
	// new linked list
	queue := list.New()

	fmt.Println("Queue:")
	// Simply append to enqueue.
	queue.PushBack(10)
	queue.PushBack(20)
	queue.PushBack(30)

	// Dequeue
	for queue.Len() > 0 {
		front := queue.Front()
		fmt.Println(front.Value)
		queue.Remove(front)
	}

	queue.PushBack(10)
	queue.PushBack(20)
	queue.PushBack(30)

	fmt.Println("\nStack:")

	// Stack
	for queue.Len() > 0 {
		back := queue.Back()
		fmt.Println(back.Value)
		queue.Remove(back)
	}

	fmt.Println("\nPriority Queue:")
	//Priority Queue
	nums := []int{3, 2, 20, 5, 3, 1, 2, 5, 6, 9, 10, 4}

	// initialize the heap data structure
	h := &IntHeap{}

	// add all the values to heap, O(n log n)
	for _, val := range nums { // O(n)
		heap.Push(h, val) // O(log n)
	}

	// print all the values from the heap
	// which should be in ascending order
	for i := 0; i < len(nums); i++ {
		fmt.Println(heap.Pop(h).(int))
	}
}
