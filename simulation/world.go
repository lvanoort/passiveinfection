package simulation

import (
	"fmt"
	"github.com/lvanoort/passiveinfection/event"
	"io"
)

// The World manages the state of the overall
// simulation and current global state of the system
type World interface {
	// Retrieve the number of timesteps this World has executed
	TimestepCount() int

	// Performs a single timestep in the simulation
	// what this timestep actually entails depends on the
	// Agents and Reservoirs in use
	ExecuteTimestep()

	// Retrieve all reservoirs in the system
	// implementations should be written so this returns
	// the same slice throughout the simulation. Removing
	// Agents from the simulation should be done in Agent
	// or Reservoir logic
	GetReservoirs() []Reservoir

	// Retrieve all the agents in the simulation
	// implementations should be written so this returns
	// the same slice throughout the simulation. Removing
	// Agents from the simulation should be done in Agent
	// or Reservoir logic
	GetAgents() []Agent

	// Retrieve all unique events present in the system.
	GetAllEvents() map[event.EventId]event.Event

	// Records an event. All reservoirs should call this
	// when generating events
	RecordEvent(event.Event)
}

func PrintShortSummary(w io.Writer, world World) {
	fmt.Fprintf(w, "Timestep %d: %d events\n", world.TimestepCount(), len(world.GetAllEvents()))
}

func PrintWorldSummary(w io.Writer, world World) {
	events := world.GetAllEvents()
	eventTotal := len(events)
	reses := world.GetReservoirs()
	agents := world.GetAgents()

	fmt.Fprintf(w, "World Summary:\n%d Timesteps\n%d Events\n%d Reservoirs\n%d Agents\n\n", world.TimestepCount(), eventTotal, len(reses), len(agents))
	for i, res := range reses {
		resEventTotal := len(res.RetrieveEvents())
		percentage := 100. * (float64(resEventTotal) / float64(eventTotal))
		fmt.Fprintf(w, "Reservoir %d: %.2f%% events known\n", i, percentage)
	}
}
