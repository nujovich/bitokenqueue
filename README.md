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
Executing '1600175640538:Event2' Done processing: Event2
```
The callback function inside each event is being called as a go routine. This was thought in case of more than one event sharing the same priority. To see the results of the go routine before giving the control back to main() function, I added a sleep of 1 second to see the printed message: Executing '1600175640538:Event2'
```bash
go e.Callback(e.Priority, e.Data)
time.Sleep(1 * time.Second)
```

2. The main() function contains a set of pre-defined events to initialize the priority queue and to show the update() and delete() functions. Feel free to change this information to suit your need

3. By default, the application doesn't generate log information. To change that, update the env variable Logs located in env.json to true
```bash
{
    "Logs": true
}
```

## Logs file
The logs file will be located at root directory of the application. The file is called mylogger.log
```bash
INFO: main.go:73: Done processing: Event2
INFO: main.go:73: Done processing: Event1
INFO: main.go:81: Waiting all go routines to finish
INFO: main.go:82: Done processing queue, elements 0
```

## Go documentation
As part of the source code

## Time-space analysis - Concurrency and free resources
I decided to implement a while iteration until the priority queue's length arrives to zero. During the loop I get the first element of the heap, knowing that's the event with top priority for execution. I analize if this event needs to be executed at this time plus an interval of 1 second. If this happends to be true, the callback function is being invoked as a go routine. I decided also to implement a wait group variable. When the go routine is invoked the wg counter increments by one. During the execution, the countiner is decreased. When all the events are executed and the Len() returns zero, I invoke the Wait() method to free resources from the go routines in case there's any being taken by that time
```bash
//At the beggining, initialized wg variable
wg := new(sync.WaitGroup)
.
.
.
.
//Add 1 to wg variable every time callback is invoked
wg.Add(1)
go e.Callback(e.Priority, e.Data)
time.Sleep(1 * time.Second)
.
.
.
.
//Notify that the go routine is finished inside callback
defer wg.Done()
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