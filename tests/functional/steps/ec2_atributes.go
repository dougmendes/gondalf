package steps

import (
	"context"
	"fmt"

	"github.com/cucumber/godog"
	"github.com/dougmendes/gondalf/tests/functional/drivers"
)

func ARegion(ctx context.Context, instanceRegion string) (context.Context, error) {
	childCtx := context.WithValue(ctx, instanceRegionKey, instanceRegion)
	return childCtx, nil
}

func TheUserRequestAllInstancesFromThatRegion(ctx context.Context) (context.Context, error) {
	instanceRegion := ctx.Value(instanceRegionKey).(string)
	api := ctx.Value(DriverContextKey).(*drivers.APIDriver)
	result, err := api.GetAllInstances(instanceRegion)

	if err != nil {
		return ctx, fmt.Errorf("error getting a list of instances for region %s: %s", instanceRegion, err)
	}

	return context.WithValue(ctx, listOfInstancesKey, result), nil
}

func RespondWithAListOfAllInstances(arg1 int) error {
	return godog.ErrPending
}
