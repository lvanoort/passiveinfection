package simulation

import (
	"github.com/lvanoort/passiveinfection/event"
)

// A Reservoir acts as a source and repository of events
// The evcents are then passively communicated through
// the network via infecting encountered Agents with
// a portion of the event log
type Reservoir interface {
	// Handle updating passive agent event list
	// and creating any events necessary
	EncounterAgent(agent Agent, world World)

	// Find out all the events this reservoir
	// knows about
	RetrieveEvents() []event.Event
}

// Create a new reservoir. This is used in world
// creation, please do not depend on the World's
// []Agent collection as Agent/Reservoir
// initialization may occur in any order
type ReservoirFactory func(World) Reservoir
