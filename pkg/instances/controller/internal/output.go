package internal

import "github.com/dougmendes/gondalf/pkg/instances/model"

func MapInstancesListToInstancesListResponse(inst *model.InstanceList) InstanceListResponse {
	result := InstanceListResponse{
		Instances: make([]InstanceResponse, 0),
	}
	for _, i := range inst.Instance {
		result.Instances = append(result.Instances, InstanceResponse{
			Name:   i.Name,
			Region: i.Region,
			Family: i.Family,
		})
	}

	return result
}

type InstanceListResponse struct {
	Instances []InstanceResponse
}

type InstanceResponse struct {
	Name   string `json:"name"`
	Region string `json:"region"`
	Family string `json:"family"`
}
