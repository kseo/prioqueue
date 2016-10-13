package prioqueue

import "errors"

var (
	// ErrQueueEmpty is an error value that is returned when the queue is empty.
	ErrQueueEmpty = errors.New("Queue is empty")
)

// CmpFunc is a function comparing two elements. The function returns a positive value if a > b, a negative value if a < b, and 0 otherwise.
type CmpFunc func(a, b interface{}) int

// PriorityQueue is a binary-heap based priority queue.
type PriorityQueue struct {
	elements []interface{}
	cmp      CmpFunc
}

// NewPriorityQueue creates a PriorityQueue from elements.
func NewPriorityQueue(elements []interface{}, cmp CmpFunc) *PriorityQueue {
	pq := &PriorityQueue{elements: elements, cmp: cmp}
	pq.buildHeap()
	return pq
}

// Add adds value to this queue.
// Worst case complexity is O(log length).
func (pq *PriorityQueue) Add(value interface{}) {
	pq.elements = append(pq.elements, value)
	last := len(pq.elements) - 1
	pq.bubbleUp(last)
}

// Peek returns the maximum element of this queue without removing it.
// Worst case complexity is O(1).
func (pq *PriorityQueue) Peek() (interface{}, error) {
	if pq.IsEmpty() {
		return nil, ErrQueueEmpty
	}
	return pq.elements[0], nil
}

// RemoveMax returns the maximum element from this queue and returns it.
// Worst case complexity is O(log length).
func (pq *PriorityQueue) RemoveMax() (interface{}, error) {
	if pq.IsEmpty() {
		return nil, ErrQueueEmpty
	}
	last := len(pq.elements) - 1
	oldRoot := pq.elements[0]
	oldLast := pq.elements[last]
	pq.elements[0] = oldLast
	pq.elements = pq.elements[:last]
	pq.heapify(0)
	return oldRoot, nil
}

// Remove removes the first occurence of value from this queue.
// Returns true if value was in the queue, false otherwise.
// Worst case complexity is O(length).
func (pq *PriorityQueue) Remove(value interface{}) bool {
	i := 0
	for i < len(pq.elements) && pq.elements[i] != value {
		i++
	}
	if i != len(pq.elements) {
		pq.removeAt(i)
		return true
	}
	return false
}

func (pq *PriorityQueue) removeAt(i int) {
	pq.forceBubbleUp(i)
	pq.RemoveMax()
}

// IsEmpty returns true if this queue is empty.
func (pq *PriorityQueue) IsEmpty() bool {
	return len(pq.elements) == 0
}

// Len returns the number of elements contained in this queue.
func (pq *PriorityQueue) Len() int {
	return len(pq.elements)
}

func (pq *PriorityQueue) buildHeap() {
	for i := len(pq.elements) / 2; i >= 0; i-- {
		pq.heapify(i)
	}
}

func (pq *PriorityQueue) heapify(i int) {
	left := 2*i + 1
	right := 2*i + 2
	largest := i

	if left < len(pq.elements) && pq.cmp(pq.elements[left], pq.elements[largest]) > 0 {
		largest = left
	}
	if right < len(pq.elements) && pq.cmp(pq.elements[right], pq.elements[largest]) > 0 {
		largest = right
	}
	if largest != i {
		pq.swap(i, largest)
		pq.heapify(largest)
	}
}

func (pq *PriorityQueue) swap(i, j int) {
	temp := pq.elements[i]
	pq.elements[i] = pq.elements[j]
	pq.elements[j] = temp
}

func (pq *PriorityQueue) bubbleUp(i int) {
	if i == 0 {
		return
	}
	parent := (i - 1) / 2
	if pq.cmp(pq.elements[parent], pq.elements[i]) < 0 {
		pq.swap(i, parent)
		pq.bubbleUp(parent)
	}
}

func (pq *PriorityQueue) forceBubbleUp(i int) {
	if i == 0 {
		return
	}
	parent := (i - 1) / 2
	pq.swap(i, parent)
	pq.forceBubbleUp(parent)
}
