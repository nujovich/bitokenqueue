package main

import (
	"container/heap"

	"github.com/nujovich/bitokenqueue/queue"

	"github.com/nujovich/bitokenqueue/event"
)

func main() {
	// Some items and their priorities.
	events := map[string]int64{
		"Event1": 1600092000, "Event2": 1600192300, "Event3": 1600092300,
	}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(queue.PriorityQueue, len(events))
	i := 0
	for data, priority := range events {
		pq[i] = &event.Event{
			Data:     data,
			Priority: priority,
			Index:    i,
		}
		i++
	}
	heap.Init(&pq)

	// Insert a new item and then modify its priority.
	e := &event.Event{
		Data:     "Event4",
		Priority: 1600091598,
	}
	heap.Push(&pq, e)
	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		e := heap.Pop(&pq).(*event.Event)
		e.Callback(e.Priority, e.Data)
	}
}
