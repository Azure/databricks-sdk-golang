// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package azure

import (
	"encoding/json"
	"net/http"

	"github.com/Azure/databricks-sdk-golang/azure/libraries/httpmodels"
)

// LibrariesAPI exposes the Libraries API
type LibrariesAPI struct {
	Client DBClient
}

func (a LibrariesAPI) init(client DBClient) LibrariesAPI {
	a.Client = client
	return a
}

// AllClusterStatuses gets the status of all libraries on all clusters
func (a LibrariesAPI) AllClusterStatuses() (httpmodels.AllClusterStatusesResp, error) {
	var resp httpmodels.AllClusterStatusesResp

	jsonResp, err := a.Client.performQuery(http.MethodGet, "/libraries/all-cluster-statuses", nil, nil)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(jsonResp, &resp)
	return resp, err
}

// ClusterStatus get the status of libraries on a cluster
func (a LibrariesAPI) ClusterStatus(req httpmodels.ClusterStatusReq) (httpmodels.ClusterStatusResp, error) {
	var resp httpmodels.ClusterStatusResp

	jsonResp, err := a.Client.performQuery(http.MethodGet, "/libraries/cluster-status", req, nil)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(jsonResp, &resp)
	return resp, err
}

// Install installs libraries on a cluster
func (a LibrariesAPI) Install(req httpmodels.InstallReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/libraries/install", req, nil)
	return err
}

// Uninstall sets libraries to be uninstalled on a cluster
func (a LibrariesAPI) Uninstall(req httpmodels.UninstallReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/libraries/uninstall", req, nil)
	return err
}
