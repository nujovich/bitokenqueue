package queue

import "github.com/nujovich/bitokenqueue/event"

type EventQueue struct {
	Head   *event.Event
	Tail   *event.Event
	Length int
}
