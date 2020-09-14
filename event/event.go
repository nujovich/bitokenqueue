package event

import "fmt"

type Event struct {
	Priority int64  //the date that the event must run
	Data     string //the atbiratry data of the item
	Index    int    //the index of the item in the heap
}

func (e Event) Callback(priority int64, data string) {
	fmt.Printf("Executing '%.2d:%s' ", priority, data)
}
