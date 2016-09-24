package event

type basicEvent struct {
	id EventId
}

// Creates an empty event with a new random ID
func GenerateBasicEvent() Event {
	return &basicEvent{id: GenerateRandomEventId()}
}

func (e *basicEvent) Id() EventId {
	return e.id
}
