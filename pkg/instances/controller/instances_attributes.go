package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	. "github.com/dougmendes/gondalf/pkg/infra/controller"
	"github.com/dougmendes/gondalf/pkg/instances/controller/internal"
	"github.com/dougmendes/gondalf/pkg/instances/usecases"
)

func NewInstanceController(solver usecases.InstanceSolver) *InstanceController {
	return &InstanceController{solver}
}

type InstanceController struct {
	solver usecases.InstanceSolver
}

func (c *InstanceController) AddRoutes(router Router) {
	router.Get("/instances/{region}", c.getAllInstances)
}

func (c *InstanceController) getAllInstances(w http.ResponseWriter, r *http.Request) error {
	params := Params(r)
	region, _ := params.String("region")
	instancesList, _ := c.solver.GetAllInstances(region)

	response := internal.MapInstancesListToInstancesListResponse(instancesList)
	return EncodedJSON(w, response, http.StatusOK)
}

// web
type uriParamsContextKey struct{}
type URIParams map[string]string

func Params(r *http.Request) URIParams {
	if v, ok := r.Context().Value(uriParamsContextKey{}).(URIParams); ok {
		return v
	}

	return nil
}

func (p URIParams) String(param string) (string, error) {
	v, ok := p[param]
	if !ok {
		err := fmt.Sprintf("%v uri param is not found: %s", http.StatusInternalServerError, param)
		return "", errors.New(err)
	}
	return v, nil
}
func EncodedJSON(w http.ResponseWriter, v interface{}, code int) error {
	if headerer, ok := v.(http.Header); ok {
		for k, values := range headerer {
			for _, v := range values {
				w.Header().Add(k, v)
			}
		}
	}

	var jsonData []byte
	var err error
	switch v := v.(type) {
	case []byte:
		jsonData = v
	case io.Reader:
		jsonData, err = ioutil.ReadAll(v)
	default:
		jsonData, err = json.Marshal(v)
	}

	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	w.WriteHeader(code)
	if _, err := w.Write(jsonData); err != nil {
		return err
	}
	return nil
}
