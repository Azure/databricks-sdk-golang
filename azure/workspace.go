package azure

import (
	"encoding/json"
	"net/http"

	"github.com/polar-rams/databricks-sdk-golang/azure/workspace/httpmodels"
)

// WorkspaceAPI exposes the Workspace API
type WorkspaceAPI struct {
	Client DBClient
}

func (a WorkspaceAPI) init(client DBClient) WorkspaceAPI {
	a.Client = client
	return a
}

// Delete an object or a directory (and optionally recursively deletes all objects in the directory)
func (a WorkspaceAPI) Delete(req httpmodels.DeleteReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/workspace/delete", req, nil)
	return err
}

// Export a notebook or contents of an entire directory
func (a WorkspaceAPI) Export(req httpmodels.ExportReq) (httpmodels.ExportResp, error) {
	var resp httpmodels.ExportResp

	jsonResp, err := a.Client.performQuery(http.MethodGet, "/workspace/export", req, nil)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(jsonResp, &resp)
	return resp, err
}

// Gets the status of an object or a directory
func (a WorkspaceAPI) GetStatus(req httpmodels.GetStatusReq) (httpmodels.GetStatusResp, error) {
	var resp httpmodels.GetStatusResp

	jsonResp, err := a.Client.performQuery(http.MethodGet, "/workspace/get-status", req, nil)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(jsonResp, &resp)
	return resp, err
}

// Import a notebook or the contents of an entire directory
func (a WorkspaceAPI) Import(req httpmodels.ImportReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/workspace/import", req, nil)
	return err
}

// List lists the contents of a directory, or the object if it is not a directory
func (a WorkspaceAPI) List(req httpmodels.ListReq) (httpmodels.ListResp, error) {
	var resp httpmodels.ListResp

	jsonResp, err := a.Client.performQuery(http.MethodGet, "/workspace/list", req, nil)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(jsonResp, &resp)
	return resp, err
}

// Mkdirs creates the given directory and necessary parent directories if they do not exists
func (a WorkspaceAPI) Mkdirs(req httpmodels.MkdirsReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/workspace/mkdirs", req, nil)
	return err
}
