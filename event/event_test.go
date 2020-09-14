package event

import "testing"

func TestEvent_Callback(t *testing.T) {
	e1 := Event{
		Priority: 1600192960000,
		Data:     "Event1",
		Index:    0,
	}
	type args struct {
		priority int64
		data     string
	}
	tests := []struct {
		name string
		e    Event
		args args
	}{
		{"Testing callback func()", e1, args{priority: 1600192960000, data: "Event1"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.e.Callback(tt.args.priority, tt.args.data)
		})
	}
}
