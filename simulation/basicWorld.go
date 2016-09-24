package simulation

import (
	"github.com/lvanoort/passiveinfection/event"
	"sync"
)

type basicWorld struct {
	agents     []Agent
	reservoirs []Reservoir
	events     map[event.EventId]event.Event
	eventLock  sync.RWMutex
	timestepCount int
}

func CreateBasicWorld(agentCount int, agentFact AgentFactory, reservoirCount int, resFact ReservoirFactory) World {
	world := &basicWorld{agents: make([]Agent, 0, agentCount), reservoirs: make([]Reservoir, 0, reservoirCount),
		events: make(map[event.EventId]event.Event)}


	for i := 0; i < agentCount; i++ {
		world.agents = append(world.agents, agentFact(world))
	}

	for i := 0; i < reservoirCount; i++ {
		world.reservoirs = append(world.reservoirs, resFact(world))
	}

	return world
}

func (b *basicWorld) TimestepCount() int {
	return b.timestepCount
}

func (b *basicWorld) ExecuteTimestep() {
	for _, agent := range b.agents {
		agent.ExecuteTimestep(b)
	}
	b.timestepCount++
}

func (b *basicWorld) GetReservoirs() []Reservoir {
	return b.reservoirs
}

func (b *basicWorld) GetAgents() []Agent {
	return b.agents
}

func (b *basicWorld) GetAllEvents() map[event.EventId]event.Event {
	b.eventLock.RLock()
	defer b.eventLock.RUnlock()

	// make a deep copy of the map
	events := make(map[event.EventId]event.Event)
	for k, v := range b.events {
		events[k] = v
	}

	return events
}

func (b *basicWorld) RecordEvent(e event.Event) {
	b.eventLock.Lock()
	defer b.eventLock.Unlock()

	b.events[e.Id()] = e
}
