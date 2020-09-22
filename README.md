# Bitoken Excercise - Implementation of a Priority Queue

This exercise consists of generating a Priority Queue that implements the heap Golang interface from /container/heap module. The priority queue contains a set of events. Each event has a date in which it will be executed and a string data containing the description/name. When the event is eligible for execution, a Callback function will prompt a message to show the user which event is being executed at that moment. The priority queue implements the following heap methods:

1. Len()
2. Swap()
3. Push()
4. Pop()
5. Update()

## Stack used
1. Vs Studio Code as IDE
2. Golang 1.15

## Go modules

The application runs under Go modules. Inspect the go.mod file for further information

```bash
go mod init github.com/nujovich/bitokenqueue
```

## Installation
1. Proceed to download the [GIT](https://github.com/nujovich/bitokenqueue.git) repository

## Execution
1. Once in your ide, open a terminal and write
```bash
go run main.go
```
When the events execution starts you will see the information about the event being executed:
```bash
Event : Event3, 2020-09-21T16:22:55.7114467-03:00
Event : Event2, 2020-09-21T16:24:55.7114467-03:00
Event : Event1, 2020-09-21T16:26:55.7114467-03:00
```
The callback function inside each event is being called as a go routine. This was thought in case of more than one event sharing the same priority. 
```bash
go e.Callback(e.Priority, e.Data, wg)
```

2. By default, the application doesn't generate log information. To change that, update the env variable Logs located in env.json to true
```bash
{
    "Logs": true
}
```

## Logs file
The logs file will be located at root directory of the application. The file is called mylogger.log
```bash
INFO: main.go:75: Done processing: Event3 Date Priority: 2020-09-21T16:16:16.7396926-03:00
INFO: main.go:75: Done processing: Event2 Date Priority: 2020-09-21T16:18:16.7396926-03:00
INFO: main.go:75: Done processing: Event1 Date Priority: 2020-09-21T16:20:16.7396926-03:00
INFO: main.go:82: Waiting all go routines to finish
INFO: main.go:83: Done processing queue, elements 0
```

## Go documentation
As part of the source code

## Time-space analysis - Concurrency and free resources
I decided to implement a while iteration until the priority queue's length arrives to zero. During the loop, I pop the first element of the queue and implement the Sleep() function until the event is ready for execution. When that happens, the callback function is being invoked as a go routine. I decided also to implement a wait group variable. When the go routine is invoked the wg counter increments by one. During the execution, the countiner is decreased. When all the events are executed and the Len() returns zero, I invoke the Wait() method to free resources from the go routines in case there's any being taken by that time
```bash
//At the beggining, initialized wg variable
wg := new(sync.WaitGroup)
.
.
.
.
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
.
.
.
.
//Notify that the go routine is finished inside callback
defer wg.Done()
time.Sleep(2 * time.Second)
//When the execution is finished, invoke Wait() 
//to make sure all go routines are finished and free resources
wg.Wait()
```

## Usage
The case of study can be implemented in a programmed money transaction, when every event is a transaction and the priority is given by the timestamp when the user decides to perform the operation


## Unit testing
100% coverage of main functions
```bash
C:\Users\Nujovich\Documents\BitokenQueue>go test github.com/nujovich/bitokenqueue/event -cover
ok      github.com/nujovich/bitokenqueue/event  0.305s  coverage: 100.0% of statements

C:\Users\Nujovich\Documents\BitokenQueue>go test github.com/nujovich/bitokenqueue/queue -cover
ok      github.com/nujovich/bitokenqueue/queue  0.306s  coverage: 100.0% of statements
```