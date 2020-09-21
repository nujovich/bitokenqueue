package event

import (
	"sync"
	"testing"
	"time"
)

func TestEvent_Callback(t *testing.T) {
	wg1 := new(sync.WaitGroup)
	wg1.Add(1)
	e1 := Event{
		Priority: time.Now().Add(time.Minute * 2),
		Data:     "Event1",
		Index:    0,
	}
	type args struct {
		priority time.Time
		data     string
		wg       *sync.WaitGroup
	}
	tests := []struct {
		name string
		e    Event
		args args
	}{
		{"Testing callback func()", e1, args{priority: time.Now().Add(time.Minute * 2), data: "Event1", wg: wg1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.e.Callback(tt.args.priority, tt.args.data, tt.args.wg)
		})
	}
}
