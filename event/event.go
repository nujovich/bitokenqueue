//Package event provide a struct that represent the information about an event
package event

import (
	"fmt"
	"sync"
	"time"
)

type Event struct {
	Priority time.Time //the date that the event must run
	Data     string    //the atbiratry data of the item
	Index    int       //the index of the item in the heap
}

//Callback just outputs a message when the event is
//eligible for execution on the heap
func (e Event) Callback(priority time.Time, data string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Event : " + data + ", " + priority.Format(time.RFC3339Nano))
	time.Sleep(2 * time.Second)
}
