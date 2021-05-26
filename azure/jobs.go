// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package azure

import (
	"encoding/json"
	"net/http"

	"github.com/Azure/databricks-sdk-golang/azure/jobs/httpmodels"
)

// JobsAPI exposes Jobs API endpoints
type JobsAPI struct {
	Client DBClient
}

func (a JobsAPI) init(client DBClient) JobsAPI {
	a.Client = client
	return a
}

// Create creates a new job
func (a JobsAPI) Create(req httpmodels.CreateReq) (httpmodels.CreateResp, error) {
	var resp httpmodels.CreateResp

	jsonResp, err := a.Client.performQuery(http.MethodPost, "/jobs/create", req, nil)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(jsonResp, &resp)
	return resp, err
}

// List lists all jobs
func (a JobsAPI) List() (httpmodels.ListResp, error) {
	var resp httpmodels.ListResp

	jsonResp, err := a.Client.performQuery(http.MethodGet, "/jobs/list", nil, nil)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(jsonResp, &resp)
	return resp, err
}

// Delete deletes a job by ID
func (a JobsAPI) Delete(req httpmodels.DeleteReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/jobs/delete", req, nil)
	return err
}

// Get gets a job by ID
func (a JobsAPI) Get(req httpmodels.GetReq) (httpmodels.GetResp, error) {
	var resp httpmodels.GetResp

	jsonResp, err := a.Client.performQuery(http.MethodGet, "/jobs/get", req, nil)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(jsonResp, &resp)
	return resp, err
}

// Reset overwrites job settings
func (a JobsAPI) Reset(req httpmodels.ResetReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/jobs/reset", req, nil)
	return err
}

// Update adds, changes, or removes specific settings of an existing job
func (a JobsAPI) Update(req httpmodels.UpdateReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/jobs/update", req, nil)
	return err
}

// RunNow runs a job now and return the run_id of the triggered run
func (a JobsAPI) RunNow(req httpmodels.RunNowReq) (httpmodels.RunNowResp, error) {
	var resp httpmodels.RunNowResp

	jsonResp, err := a.Client.performQuery(http.MethodPost, "/jobs/run-now", req, nil)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(jsonResp, &resp)
	return resp, err
}

// RunsSubmit submit a one-time run
func (a JobsAPI) RunsSubmit(req httpmodels.RunsSubmitReq) (httpmodels.RunsSubmitResp, error) {
	var resp httpmodels.RunsSubmitResp

	jsonResp, err := a.Client.performQuery(http.MethodPost, "/jobs/runs/submit", req, nil)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(jsonResp, &resp)
	return resp, err
}

// RunsList lists runs from most recently started to least
func (a JobsAPI) RunsList(req httpmodels.RunsListReq) (httpmodels.RunsListResp, error) {
	var resp httpmodels.RunsListResp

	jsonResp, err := a.Client.performQuery(http.MethodGet, "/jobs/runs/list", req, nil)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(jsonResp, &resp)
	return resp, err
}

// RunsGet retrieve the metadata of a run
func (a JobsAPI) RunsGet(req httpmodels.RunsGetReq) (httpmodels.RunsGetResp, error) {
	var resp httpmodels.RunsGetResp

	jsonResp, err := a.Client.performQuery(http.MethodGet, "/jobs/runs/get", req, nil)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(jsonResp, &resp)
	return resp, err
}

// RunsExport exports and retrieve the job run task
func (a JobsAPI) RunsExport(req httpmodels.RunsExportReq) (httpmodels.RunsExportResp, error) {
	var resp httpmodels.RunsExportResp

	jsonResp, err := a.Client.performQuery(http.MethodGet, "/jobs/runs/export", req, nil)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(jsonResp, &resp)
	return resp, err
}

// RunsCancel cancels a run
func (a JobsAPI) RunsCancel(req httpmodels.RunsCancelReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/jobs/runs/cancel", req, nil)
	return err
}

// RunsGetOutput retrieves the output of a run
func (a JobsAPI) RunsGetOutput(req httpmodels.RunsGetOutputReq) (httpmodels.RunsGetOutputResp, error) {
	var resp httpmodels.RunsGetOutputResp

	jsonResp, err := a.Client.performQuery(http.MethodGet, "/jobs/runs/get-output", req, nil)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(jsonResp, &resp)
	return resp, err
}

// RunsDelete deletes a non-active run. Returns an error if the run is active.
func (a JobsAPI) RunsDelete(req httpmodels.RunsDeleteReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/jobs/runs/delete", req, nil)
	return err
}
