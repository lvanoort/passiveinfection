package simulation

import (
	"github.com/lvanoort/passiveinfection/event"
	"math/rand"
	"sync"
)

// Visit probability in hundred thousandths
// eg. a value of 1000 would represent 1000/100000 odds
// or 1% chance of the agent visiting a reservoir in each timestep
type VisitProbability int

var DefaultVisitProbability int = 5000
var DefaultEventCapacity int = 20

type BasicAgent struct {
	events           map[event.EventId]event.Event
	capacity         int
	visitProbability int
	eventMutex       sync.RWMutex
}

func BasicAgentFact(w World) Agent {
	return NewBasicAgent()
}

func NewBasicAgent() Agent {
	return &BasicAgent{events: make(map[event.EventId]event.Event), capacity: DefaultEventCapacity, visitProbability: DefaultVisitProbability}
}

func (b *BasicAgent) RetrieveEvents() []event.Event {
	b.eventMutex.RLock()
	defer b.eventMutex.RUnlock()

	events := make([]event.Event, 0, len(b.events))
	for _, v := range b.events {
		events = append(events, v)
	}

	return events
}

func (b *BasicAgent) GetMaxEventCount() int {
	return b.capacity
}

func (b *BasicAgent) PutEvents(events []event.Event) {
	b.eventMutex.Lock()
	defer b.eventMutex.Unlock()

	b.events = make(map[event.EventId]event.Event)
	for i := 0; i < len(events) && i < b.capacity; i++ {
		event := events[i]
		b.events[event.Id()] = event
	}
}

func (b *BasicAgent) ExecuteTimestep(world World) {
	roll := 1 + rand.Intn(100000)
	if roll <= b.visitProbability {
		res := world.GetReservoirs()
		chosen := res[rand.Intn(len(res))]
		chosen.EncounterAgent(b, world)
	}
}
