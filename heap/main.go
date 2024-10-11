package main

import "fmt"

func main() {
	testValue := []int{4, 3, 2, 1, 0, 9, 8, 7, 6, 5}
	heap := Heap{}
	for _, value := range testValue {
		heap.Push(value)
	}
	fmt.Println("Top: ", heap.Top())
	fmt.Println("data: ", heap.heap)
	for heap.Len() > 0 {
		fmt.Println(heap.Pop())
		fmt.Println("data: ", heap.heap)
	}
	return
}

//min heap
type Heap struct {
	heap []int
}

func (h *Heap) Len() int {
	return len(h.heap)
}

func (h *Heap) Push(value int) {
	h.heap = append(h.heap, value)
	for i := len(h.heap) - 1; i > 0; {
		if h.heap[i] < h.heap[(i-1)/2] {
			h.heap[i], h.heap[(i-1)/2] = h.heap[(i-1)/2], h.heap[i]
			i = (i - 1) / 2
		} else {
			break
		}
	}
}

func (h *Heap) Top() int {
	if len(h.heap) == 0 {
		return -1
	}
	return h.heap[0]
}

func (h *Heap) Pop() int {
	if len(h.heap) == 0 {
		return -1
	}
	value := h.heap[0]
	h.heap[0] = h.heap[len(h.heap)-1]
	h.heap = h.heap[:len(h.heap)-1]
	for i := 0; i < len(h.heap); {
		if 2*i+2 < len(h.heap) && h.heap[i] > h.heap[2*i+1] && h.heap[i] > h.heap[2*i+2] {
			if h.heap[2*i+1] < h.heap[2*i+2] {
				h.heap[i], h.heap[2*i+1] = h.heap[2*i+1], h.heap[i]
				i = 2*i + 1
			} else {
				h.heap[i], h.heap[2*i+2] = h.heap[2*i+2], h.heap[i]
				i = 2*i + 2
			}
		} else if 2*i+1 < len(h.heap) && h.heap[i] > h.heap[2*i+1] {
			h.heap[i], h.heap[2*i+1] = h.heap[2*i+1], h.heap[i]
			i = 2*i + 1
		} else if 2*i+2 < len(h.heap) && h.heap[i] > h.heap[2*i+2] {
			h.heap[i], h.heap[2*i+2] = h.heap[2*i+2], h.heap[i]
			i = 2*i + 2
		} else {
			break
		}
	}

	return value
}
