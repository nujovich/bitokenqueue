//queue provides a customized implementation of the heap interface
package queue

import (
	"container/heap"

	"github.com/nujovich/bitokenqueue/event"
)

//Priority queue is a type that will hold an array of events pointers
type PriorityQueue []*event.Event

//Len() returns the length of the queue
func (pq PriorityQueue) Len() int {
	return len(pq)
}

//Less() takes as param two indexes and compares the priority on both events
//It returns true if first event will execute as top priority
//It return false on the contrary
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

//Swap() takes as param two indexes and changes the information holding on both events indexes
//Not being used at current implementation
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

//Push() pushes the element x onto the heap. The complexity is O(log n) where n = h.Len().
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	event := x.(*event.Event)
	event.Index = n
	*pq = append(*pq, event)
}

//Pop() removes and returns the minimum element (according to Less) from the heap.
//The complexity is O(log n) where n = h.Len(). Pop is equivalent to Remove(h, 0).
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	event := old[n-1]
	old[n-1] = nil   // avoid memory leak
	event.Index = -1 // for safety
	*pq = old[0 : n-1]
	return event
}

//Up() modifies the priority and data of an event in the queue.
func (pq *PriorityQueue) Update(event *event.Event, value string, priority int64) {
	event.Data = value
	event.Priority = priority
	heap.Fix(pq, event.Index)
}
