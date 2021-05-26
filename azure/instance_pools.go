package azure

import (
	"encoding/json"
	"net/http"

	"github.com/polar-rams/databricks-sdk-golang/azure/instance_pools/httpmodels"
)

// InstancePoolsAPI exposes the InstancePools API
type InstancePoolsAPI struct {
	Client DBClient
}

func (a InstancePoolsAPI) init(client DBClient) InstancePoolsAPI {
	a.Client = client
	return a
}

// Create creates an instance pool
func (a InstancePoolsAPI) Create(req httpmodels.CreateReq) (httpmodels.CreateResp, error) {
	var resp httpmodels.CreateResp

	jsonResp, err := a.Client.performQuery(http.MethodPost, "/instance-pools/create", req, nil)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(jsonResp, &resp)
	return resp, err
}

// Edit modifies the configuration of an existing instance pool.
func (a InstancePoolsAPI) Edit(req httpmodels.EditReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/instance-pools/edit", req, nil)
	return err
}

// Delete permanently deletes the instance pool.
func (a InstancePoolsAPI) Delete(req httpmodels.DeleteReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/instance-pools/delete", req, nil)
	return err
}

// Get retrieves the information for an instance pool given its identifier.
func (a InstancePoolsAPI) Get(req httpmodels.GetReq) (httpmodels.GetResp, error) {
	var resp httpmodels.GetResp

	jsonResp, err := a.Client.performQuery(http.MethodGet, "/instance-pools/get", req, nil)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(jsonResp, &resp)
	return resp, err
}

// List returns information for all instance pools.
func (a InstancePoolsAPI) List() (httpmodels.ListResp, error) {
	var resp httpmodels.ListResp

	jsonResp, err := a.Client.performQuery(http.MethodGet, "/instance-pools/list", nil, nil)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(jsonResp, &resp)
	return resp, err
}
