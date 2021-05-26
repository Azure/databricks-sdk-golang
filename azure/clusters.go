// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package azure

import (
	"encoding/json"
	"net/http"

	"github.com/Azure/databricks-sdk-golang/azure/clusters/httpmodels"
)

// ClustersAPI exposes the Clusters API
type ClustersAPI struct {
	Client DBClient
}

func (a ClustersAPI) init(client DBClient) ClustersAPI {
	a.Client = client
	return a
}

// Create creates a new Spark cluster
func (a ClustersAPI) Create(req httpmodels.CreateReq) (httpmodels.CreateResp, error) {
	var resp httpmodels.CreateResp

	jsonResp, err := a.Client.performQuery(http.MethodPost, "/clusters/create", req, nil)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(jsonResp, &resp)
	return resp, err
}

// Edit edits the configuration of a cluster to match the provided attributes and size
func (a ClustersAPI) Edit(req httpmodels.EditReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/edit", req, nil)
	return err
}

// Start starts a terminated Spark cluster given its ID
func (a ClustersAPI) Start(req httpmodels.StartReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/start", req, nil)
	return err
}

// Restart restart a Spark cluster given its ID. If the cluster is not in a RUNNING state, nothing will happen.
func (a ClustersAPI) Restart(req httpmodels.RestartReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/restart", req, nil)
	return err
}

// Resize resizes a cluster to have a desired number of workers. This will fail unless the cluster is in a RUNNING state.
func (a ClustersAPI) Resize(req httpmodels.ResizeReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/resize", req, nil)
	return err
}

// Delete terminates a Spark cluster given its ID
func (a ClustersAPI) Delete(req httpmodels.DeleteReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/delete", req, nil)
	return err
}

// Terminate is an alias of Delete
func (a ClustersAPI) Terminate(req httpmodels.DeleteReq) error {
	return a.Delete(req)
}

// PermanentDelete permanently delete a cluster
func (a ClustersAPI) PermanentDelete(req httpmodels.PermanentDeleteReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/permanent-delete", req, nil)
	return err
}

// Get retrieves the information for a cluster given its identifier
func (a ClustersAPI) Get(req httpmodels.GetReq) (httpmodels.GetResp, error) {
	var resp httpmodels.GetResp

	jsonResp, err := a.Client.performQuery(http.MethodGet, "/clusters/get", req, nil)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(jsonResp, &resp)
	return resp, err
}

// Pin ensure that an interactive cluster configuration is retained even after a cluster has been terminated for more than 30 days
func (a ClustersAPI) Pin(req httpmodels.PinReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/pin", req, nil)
	return err
}

// Unpin allows the cluster to eventually be removed from the list returned by the List API
func (a ClustersAPI) Unpin(req httpmodels.UnpinReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/clusters/unpin", req, nil)
	return err
}

// List return information about all pinned clusters, currently active clusters,
// up to 70 of the most recently terminated interactive clusters in the past 30 days,
// and up to 30 of the most recently terminated job clusters in the past 30 days
func (a ClustersAPI) List() (httpmodels.ListResp, error) {
	var resp httpmodels.ListResp

	jsonResp, err := a.Client.performQuery(http.MethodGet, "/clusters/list", nil, nil)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(jsonResp, &resp)
	return resp, err
}

// ListNodeTypes returns a list of supported Spark node types
func (a ClustersAPI) ListNodeTypes() (httpmodels.ListNodeTypesResp, error) {
	var resp httpmodels.ListNodeTypesResp

	jsonResp, err := a.Client.performQuery(http.MethodGet, "/clusters/list-node-types", nil, nil)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(jsonResp, &resp)
	return resp, err
}

// RuntimeVersions return the list of available Runtime versions
func (a ClustersAPI) RuntimeVersions() (httpmodels.RuntimeVersionsResp, error) {
	var resp httpmodels.RuntimeVersionsResp

	jsonResp, err := a.Client.performQuery(http.MethodGet, "/clusters/spark-versions", nil, nil)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(jsonResp, &resp)
	return resp, err
}

// Events retrieves a list of events about the activity of a cluster
func (a ClustersAPI) Events(req httpmodels.EventsReq) (httpmodels.EventsResp, error) {
	var resp httpmodels.EventsResp

	jsonResp, err := a.Client.performQuery(http.MethodPost, "/clusters/events", req, nil)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(jsonResp, &resp)
	return resp, err
}
