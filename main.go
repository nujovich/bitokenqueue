//Current main program
package main

import (
	"container/heap"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"time"

	"github.com/nujovich/bitokenqueue/customlogger"
	"github.com/nujovich/bitokenqueue/event"
	"github.com/nujovich/bitokenqueue/queue"
)

//Configuration struct holds the env variable info setup in env.json
//By default Logs is set to false
type Configuration struct {
	Logs bool `json:"Logs"`
}

//Const holding the path to env file
const FILENAME string = "env.json"

func main() {
	jsonFile, _ := os.Open(FILENAME)
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var config Configuration
	json.Unmarshal(byteValue, &config)
	Logger, error := isLoggingTrue(config.Logs)

	// Some events and their priorities.
	wg := new(sync.WaitGroup)
	priority1 := time.Now().Add(time.Minute * 3)
	priority2 := time.Now().Add(time.Minute * 5)
	events := map[string]time.Time{
		"Event1": priority2, "Event2": priority1,
	}

	// Create a priority queue, put the events in it, and
	// establish the priority queue (heap) invariants.
	pq := make(queue.PriorityQueue, 0)
	heap.Init(&pq)
	i := 0
	for data, priority := range events {
		e1 := &event.Event{
			Data:     data,
			Priority: priority,
			Index:    i,
		}
		heap.Push(&pq, e1)
		i++
	}

	//Insert a new item and then modify its priority.
	e := &event.Event{
		Data:     "Event3",
		Priority: time.Now().Add(time.Minute * 2),
	}
	newPriority := time.Now().Add(time.Minute * 1)
	heap.Push(&pq, e)
	pq.Update(e, e.Data, newPriority)

	// Take the items out; they arrive in increasing priority order.
	for pq.Len() > 0 {
		wg.Add(1)
		e := heap.Pop(&pq).(*event.Event)
		diff := time.Until(e.Priority)
		time.Sleep(1 * diff)
		go e.Callback(e.Priority, e.Data, wg)
		if error == nil {
			Logger.Print("Done processing: " + e.Data + " Date Priority: " + e.Priority.Format(time.RFC3339Nano))
		} else {
			fmt.Println("Done processing: " + e.Data)
		}
	}
	wg.Wait()
	if error == nil {
		Logger.Print("Waiting all go routines to finish")
		Logger.Print("Done processing queue, elements ", pq.Len())
	} else {
		fmt.Println("Waiting all go routines to finish")
		fmt.Println("Done processing queue, elements ", pq.Len())
	}
}

//isLogginTrue returns an instace of CustomLogger if the Logs env variable is set to true
//it returns an error on the contrary
func isLoggingTrue(env bool) (customlogger.CustomLogger, error) {
	if env == true {
		return customlogger.GetInstance(), nil
	} else {
		return customlogger.CustomLogger{}, errors.New("No logging set")
	}
}
