package steps

type contextKey int

const (
	DriverContextKey contextKey = iota
	instanceRegionKey
	listOfInstancesKey
)
