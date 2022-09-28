package drivers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func NewAPIDriver(baseURL string) *APIDriver {
	return &APIDriver{baseURL}
}

type APIDriver struct {
	baseURL string
}

func (d *APIDriver) doGetRequest(path string) (*http.Response, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s%s", d.baseURL, path)
	req, _ := http.NewRequest("GET", url, nil)
	return client.Do(req)
}

func (d *APIDriver) GetAllInstances(region string) (InstanceList, error) {
	path := fmt.Sprintf("/instances/%s", region)
	res, err := d.doGetRequest(path)

	if err != nil {
		return InstanceList{}, fmt.Errorf("error getting information from api: %w", err)
	}
	var result InstanceList
	json.NewDecoder(res.Body).Decode(&result)
	return result, nil
}

type InstanceList struct {
	Instances []Instance `json:instance_list`
}

type Instance struct {
	name   string `json:name`
	region string `json:region`
	family string `json:family`
}
