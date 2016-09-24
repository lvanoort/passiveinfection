package simulation

import (
	"github.com/lvanoort/passiveinfection/event"
)

// Represents an infection carrier for transferring
// event data
type Agent interface {
	// Retrieves all the events this Agent is currently storing
	RetrieveEvents() []event.Event
	// Retrieves the maximum number of events that can be stored
	// in this Agent
	GetMaxEventCount() int
	// Set the events that are stored internally. If len(events)
	// is greater than GetMaxEventCount(), the Agent may truncate
	// the event set however it wishes. Additionally, if len(events)
	// is less than GetMaxEvenCount(), the Agent may replace its
	// entire event set or merge them, depending on implementation
	PutEvents(events []event.Event)

	// Perform a timestep of the simulation within the world
	ExecuteTimestep(world World)
}

// Create a new reservoir. This is used in world
// creation, please do not depend on the World's
// []Reservoir collection as Agent/Reservoir
// initialization may occur in any order
type AgentFactory func(World) Agent
