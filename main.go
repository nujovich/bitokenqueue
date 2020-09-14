package main

import (
	"container/heap"
	"fmt"
	"sync"
	"time"

	"github.com/nujovich/bitokenqueue/event"
	"github.com/nujovich/bitokenqueue/queue"
)

func main() {
	// Some events and their priorities.
	wg := new(sync.WaitGroup)
	priority1 := time.Now().Add(time.Minute*3).UnixNano() / int64(time.Millisecond)
	priority2 := time.Now().Add(time.Minute*5).UnixNano() / int64(time.Millisecond)
	events := map[string]int64{
		"Event1": priority2, "Event2": priority1,
	}

	// Create a priority queue, put the events in it, and
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

	//Insert a new item and then modify its priority.
	e := &event.Event{
		Data:     "Event3",
		Priority: time.Now().Add(time.Minute*2).UnixNano() / int64(time.Millisecond),
	}
	heap.Push(&pq, e)
	newPriority := time.Now().Add(time.Minute*1).UnixNano() / int64(time.Millisecond)
	pq.Update(e, e.Data, newPriority)
	//Delete it from queue
	heap.Remove(&pq, 0)
	// Take the items out; they arrive in increasing priority order.
	for pq.Len() > 0 {
		e := pq[0]
		now := makeTimestamp()
		interval := time.Now().Add(time.Second*1).UnixNano() / int64(time.Millisecond)
		if e.Priority >= now && e.Priority <= interval {
			e = heap.Pop(&pq).(*event.Event)
			go e.Callback(e.Priority, e.Data)
			time.Sleep(1 * time.Second)
			fmt.Println("Done processing: " + e.Data)
		}
	}
	wg.Wait()
	fmt.Println("Waiting all go routines to finish")
	fmt.Println("Done processing queue, elements ", pq.Len())
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
