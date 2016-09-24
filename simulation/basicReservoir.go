package simulation

import (
	"github.com/lvanoort/passiveinfection/event"
	"math/rand"
	"sync"
)

// The BasicReservoir generates an event per encounter
// and randomly issues up to n random events from its log
// to encountered agents where n is the capacity of the agent
// encountered. It will always include the event generated in
// the encounter if the Agent at least has a capacity of 1
type BasicReservoir struct {
	events     map[event.EventId]event.Event
	eventKeys  []event.EventId
	eventMutex sync.RWMutex
}

func BasicReservoirFact(world World) Reservoir {
	return NewBasicReservoir()
}

func NewBasicReservoir() Reservoir {
	return &BasicReservoir{events: make(map[event.EventId]event.Event)}
}

func (r *BasicReservoir) EncounterAgent(agent Agent, world World) {
	r.eventMutex.Lock()
	defer r.eventMutex.Unlock()

	for _, agentEvent := range agent.RetrieveEvents() {
		if _, ok := r.events[agentEvent.Id()]; !ok {
			r.eventKeys = append(r.eventKeys, agentEvent.Id())
			r.events[agentEvent.Id()] = agentEvent
		}
	}

	newEvent := event.GenerateBasicEvent()

	// record the new event everywhere
	world.RecordEvent(newEvent)
	r.events[newEvent.Id()] = newEvent
	r.eventKeys = append(r.eventKeys, newEvent.Id())

	agentCap := agent.GetMaxEventCount()

	if agentCap > 0 {
		newContents := make([]event.Event, 0, agentCap)
		newContents = append(newContents, newEvent)
		alreadyAdded := make(map[event.EventId]bool)

		for len(newContents) < agentCap {
			candidate := r.events[r.eventKeys[rand.Intn(len(r.eventKeys))]]

			if _, ok := alreadyAdded[candidate.Id()]; !ok {
				newContents = append(newContents, candidate)
			} else if len(alreadyAdded) == len(r.eventKeys) {
				break
			}
		}

		agent.PutEvents(newContents)
	}

}

func (r *BasicReservoir) RetrieveEvents() []event.Event {
	r.eventMutex.RLock()
	defer r.eventMutex.RUnlock()

	slice := make([]event.Event, 0, len(r.events))

	for _, v := range r.events {
		slice = append(slice, v)
	}

	return slice
}
