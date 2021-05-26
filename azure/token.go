// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package azure

import (
	"encoding/json"
	"net/http"

	"github.com/Azure/databricks-sdk-golang/azure/token/httpmodels"
)

// TokenAPI exposes the Token API
type TokenAPI struct {
	Client DBClient
}

func (a TokenAPI) init(client DBClient) TokenAPI {
	a.Client = client
	return a
}

// Create creates and return a token
func (a TokenAPI) Create(req httpmodels.CreateReq) (httpmodels.CreateResp, error) {
	var resp httpmodels.CreateResp

	jsonResp, err := a.Client.performQuery(http.MethodPost, "/token/create", req, nil)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(jsonResp, &resp)
	return resp, err
}

// List lists all the valid tokens for a user-workspace pair
func (a TokenAPI) List() (httpmodels.ListResp, error) {
	var resp httpmodels.ListResp

	jsonResp, err := a.Client.performQuery(http.MethodGet, "/token/list", nil, nil)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(jsonResp, &resp)
	return resp, err
}

// Revoke revokes an access token
func (a TokenAPI) Revoke(req httpmodels.DeleteReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/token/delete", req, nil)
	return err
}
