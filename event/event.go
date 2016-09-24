package event

import (
	"crypto/rand"
	"fmt"
)

// Event identifier, which should take the
// form of a globally unique id
type EventId string

// Represents an event in this system
// Note that all Events must be immutable for the CvRDT
// to operate safely
type Event interface {
	// Retrieve the even identifier associated with this event
	Id() EventId
}

// An EventSource may not have necessarily generated the events,
// but it represents anything that may contain events
type EventRepository interface {
	RetrieveEvents() []Event
}

func GenerateRandomEventId() EventId {
	uuid := make([]byte, 16)
	rand.Read(uuid)
	return EventId(fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]))
}
