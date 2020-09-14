package event

type Event struct {
	RunData int64
	Data    string
	Next    *Event
}
