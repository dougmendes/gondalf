package model

type InstanceList struct {
	Instance []Instance
}

type Instance struct {
	Name   string
	Region string
	Family string
}
