package usecases

import "github.com/dougmendes/gondalf/pkg/instances/model"

//go:generate mockgen -source=./main.go -destionation=../../../test/mock/instances/usecases.go
type InstanceSolver interface {
	GetAllInstances(region string) (*model.InstanceList, error)
}
