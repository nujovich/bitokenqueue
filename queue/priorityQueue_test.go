package queue

import (
	"reflect"
	"testing"
	"time"

	"github.com/nujovich/bitokenqueue/event"
)

func TestPriorityQueue_Len(t *testing.T) {
	tests := []struct {
		name string
		pq   PriorityQueue
		want int
	}{
		{"Testing len() pq1", make(PriorityQueue, 1), 1},
		{"Testing len() pq2", make(PriorityQueue, 2), 2},
		{"Testing len() pq3", make(PriorityQueue, 3), 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pq.Len(); got != tt.want {
				t.Errorf("PriorityQueue.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPriorityQueue_Less(t *testing.T) {
	//false
	pq1 := make(PriorityQueue, 4)
	e1 := &event.Event{
		Data:     "Event1",
		Priority: time.Now().Add(time.Minute * 2),
	}
	e2 := &event.Event{
		Data:     "Event2",
		Priority: time.Now().Add(time.Minute * 1),
	}
	e3 := &event.Event{
		Data:     "Event3",
		Priority: time.Now().Add(time.Minute * 2),
	}
	e4 := &event.Event{
		Data:     "Event4",
		Priority: time.Now().Add(time.Minute * 3),
	}
	pq1[0] = e1
	pq1[1] = e2
	//true
	pq1[2] = e3
	pq1[3] = e4
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		pq   PriorityQueue
		args args
		want bool
	}{
		{"Testing priority on indexes 0 & 1", pq1, args{i: 0, j: 1}, false},
		{"Testing priority on indexes 2 & 3", pq1, args{i: 2, j: 3}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pq.Less(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("PriorityQueue.Less() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPriorityQueue_Swap(t *testing.T) {
	pq1 := make(PriorityQueue, 4)
	e1 := &event.Event{
		Data:     "Event1",
		Priority: time.Now().Add(time.Minute * 3),
	}
	e2 := &event.Event{
		Data:     "Event2",
		Priority: time.Now().Add(time.Minute * 2),
	}
	pq1[0] = e1
	pq1[1] = e2
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name string
		pq   PriorityQueue
		args args
	}{
		{"Testing swap func()", pq1, args{i: 0, j: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pq.Swap(tt.args.i, tt.args.j)
		})
	}
}

func TestPriorityQueue_Push(t *testing.T) {
	pq1 := make(PriorityQueue, 1)
	e1 := &event.Event{
		Data:     "Event1",
		Priority: time.Now().Add(time.Minute * 2),
	}

	pq1[0] = e1
	pointer := *pq1[0]
	type args struct {
		x interface{}
	}
	tests := []struct {
		name string
		pq   *PriorityQueue
		args args
	}{
		{"Testing push func()", &pq1, args{x: &pointer}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pq.Push(tt.args.x)
		})
	}
}

func TestPriorityQueue_Pop(t *testing.T) {
	pq1 := make(PriorityQueue, 1)
	e1 := &event.Event{
		Data:     "Event1",
		Priority: time.Now().Add(time.Minute * 2),
	}

	pq1[0] = e1
	tests := []struct {
		name string
		pq   *PriorityQueue
		want interface{}
	}{
		{"Testing push func()", &pq1, e1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pq.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PriorityQueue.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPriorityQueue_Update(t *testing.T) {
	pq1 := make(PriorityQueue, 1)
	e1 := &event.Event{
		Data:     "Event1",
		Priority: time.Now().Add(time.Minute * 2),
	}
	type args struct {
		event    *event.Event
		value    string
		priority time.Time
	}
	tests := []struct {
		name string
		pq   *PriorityQueue
		args args
	}{
		{"Testing updating priority on event ()", &pq1, args{e1, "Event1", time.Now().Add(time.Minute * 1)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pq.Update(tt.args.event, tt.args.value, tt.args.priority)
		})
	}
}
