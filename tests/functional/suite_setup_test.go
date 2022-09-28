package main

import (
	"github.com/cucumber/godog"
	"github.com/dougmendes/gondalf/tests/functional/steps"
)

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^a region <([\w-_\s]+)>$`, steps.Aec2region)

	ctx.Step(`^respond with a list of all ec(\d+)$`, steps.RespondWithAListOfAllEc2)
	ctx.Step(`^the user request all ec(\d+) from that region$`, steps.TheUserRequestAllEc2FromThatRegion)
}
