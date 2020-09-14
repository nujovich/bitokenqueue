package queue

import (
	"container/heap"

	"github.com/nujovich/bitokenqueue/event"
)

type PriorityQueue []*event.Event

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func newPriorityQueue() *PriorityQueue {
	return &PriorityQueue{}
}

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].Priority > pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	event := x.(*event.Event)
	event.Index = n
	*pq = append(*pq, event)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	event := old[n-1]
	old[n-1] = nil   // avoid memory leak
	event.Index = -1 // for safety
	*pq = old[0 : n-1]
	return event
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(event *event.Event, value string, priority int64) {
	event.Data = value
	event.Priority = priority
	heap.Fix(pq, event.Index)
}
