package main

import (
	"github.com/cucumber/godog"
	"github.com/dougmendes/gondalf/tests/functional/steps"
)

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^a region <([\w-_\s]+)>$`, steps.ARegion)
	ctx.Step(`^respond with a list of all instances$`, steps.RespondWithAListOfAllInstances)
	ctx.Step(`^the user request all instances from that region$`, steps.TheUserRequestAllInstancesFromThatRegion)
}
