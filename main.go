package main

import (
	"github.com/lvanoort/passiveinfection/simulation"
	"os"
)

func main() {
	world := simulation.CreateBasicWorld(1000, simulation.BasicAgentFact, 20, simulation.BasicReservoirFact)
	simulation.PrintShortSummary(os.Stdout, world)

	for i := 0; i < 500; i++ {
		world.ExecuteTimestep()
		simulation.PrintShortSummary(os.Stdout, world)
	}

	simulation.PrintWorldSummary(os.Stdout, world)
}
