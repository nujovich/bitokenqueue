package main

import (
	"testing"
	"time"
)

func Test_makeTimestamp(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		{"Testing timestamp func()", time.Now().UnixNano() / int64(time.Millisecond)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeTimestamp(); got != tt.want {
				t.Errorf("makeTimestamp() = %v, want %v", got, tt.want)
			}
		})
	}
}
