package event

import "time"

type DomainEvent struct {
	Data    any
	Date    time.Time
	Name    string
	Version uint8
}

func NewDomainEvent(
	data any,
	name string,
) *DomainEvent {
	return &DomainEvent{
		Data:    data,
		Date:    time.Now(),
		Name:    name,
		Version: 1,
	}
}
