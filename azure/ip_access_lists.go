package azure

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Azure/databricks-sdk-golang/azure/ip_access_lists/httpmodels"
	"github.com/Azure/databricks-sdk-golang/azure/ip_access_lists/models"
)

// IPAccessListsAPI exposes the IPAccessLists API
type IPAccessListsAPI struct {
	Client DBClient
}

func (a IPAccessListsAPI) init(client DBClient) IPAccessListsAPI {
	a.Client = client
	return a
}

// CheckEnabled gets the IP access list feature status for Databricks workspace.
func (a IPAccessListsAPI) CheckEnabled() (httpmodels.CheckEnabledResp, error) {
	var resp httpmodels.CheckEnabledResp

	var req = httpmodels.CheckEnabledReq{Keys: "enableIpAccessLists"}
	jsonResp, err := a.Client.performQuery(http.MethodGet, "/workspace-conf", req, nil)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(jsonResp, &resp)
	return resp, err
}

// Set enables or disables the IP access list feature for Databricks workspace.
func (a IPAccessListsAPI) Set(flag bool) error {
	req := httpmodels.SetReq{
		EnableIpAccessLists: models.EnableIpAccessListsFalse,
	}
	if flag {
		req.EnableIpAccessLists = models.EnableIpAccessListsTrue
	}
	_, err := a.Client.performQuery(http.MethodPatch, "/workspace-conf", req, nil)

	return err
}

// Add adds an IP access list for Databricks workspace.
func (a IPAccessListsAPI) Add(req httpmodels.AddReq) (httpmodels.AddResp, error) {
	var resp httpmodels.AddResp

	jsonResp, err := a.Client.performQuery(http.MethodPost, "/ip-access-lists", req, nil)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(jsonResp, &resp)
	return resp, err
}

// GetAll gets all IP access lists for Databricks workspace.
func (a IPAccessListsAPI) GetAll() (httpmodels.GetAllResp, error) {
	var resp httpmodels.GetAllResp

	jsonResp, err := a.Client.performQuery(http.MethodGet, "/ip-access-lists", nil, nil)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(jsonResp, &resp)
	return resp, err
}

// Get gets an IP access list, specified by its list ID.
func (a IPAccessListsAPI) Get(id string) (httpmodels.GetResp, error) {
	var resp httpmodels.GetResp

	var path = fmt.Sprintf("/ip-access-lists/%s", id)
	jsonResp, err := a.Client.performQuery(http.MethodGet, path, nil, nil)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(jsonResp, &resp)
	return resp, err
}

// Replace replaces an IP access list, specified by its ID.
func (a IPAccessListsAPI) Replace(id string, req httpmodels.ReplaceReq) (httpmodels.ReplaceResp, error) {
	var resp httpmodels.ReplaceResp

	var path = fmt.Sprintf("/ip-access-lists/%s", id)
	jsonResp, err := a.Client.performQuery(http.MethodPut, path, req, nil)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(jsonResp, &resp)
	return resp, err
}

// Delete deletes an IP access list, specified by its list ID.
func (a IPAccessListsAPI) Delete(id string) error {
	var path = fmt.Sprintf("/ip-access-lists/%s", id)
	_, err := a.Client.performQuery(http.MethodDelete, path, nil, nil)
	return err
}

// Update modifies an existing IP access list, specified by its ID.
func (a IPAccessListsAPI) Update(id string, req httpmodels.UpdateReq) (httpmodels.UpdateResp, error) {
	var resp httpmodels.UpdateResp

	var path = fmt.Sprintf("/ip-access-lists/%s", id)
	jsonResp, err := a.Client.performQuery(http.MethodPatch, path, req, nil)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(jsonResp, &resp)
	return resp, err
}
