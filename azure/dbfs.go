package azure

import (
	"encoding/json"
	"net/http"

	"github.com/polar-rams/databricks-sdk-golang/azure/dbfs/httpmodels"
)

// DbfsAPI exposes the DBFS API
type DbfsAPI struct {
	Client DBClient
}

func (a DbfsAPI) init(client DBClient) DbfsAPI {
	a.Client = client
	return a
}

// AddBlock appends a block of data to the stream specified by the input handle
func (a DbfsAPI) AddBlock(req httpmodels.AddBlockReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/dbfs/add-block", req, nil)
	return err
}

// Close closes the stream specified by the input handle
func (a DbfsAPI) Close(req httpmodels.CloseReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/dbfs/close", req, nil)
	return err
}

// Create opens a stream to write to a file and returns a handle to this stream
func (a DbfsAPI) Create(req httpmodels.CreateReq) (httpmodels.CreateResp, error) {
	var resp httpmodels.CreateResp

	jsonResp, err := a.Client.performQuery(http.MethodPost, "/dbfs/create", req, nil)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(jsonResp, &resp)
	return resp, err
}

// Delete deletes the file or directory (optionally recursively delete all files in the directory)
func (a DbfsAPI) Delete(req httpmodels.DeleteReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/dbfs/delete", req, nil)
	return err
}

// GetStatus gets the file information of a file or directory
func (a DbfsAPI) GetStatus(req httpmodels.GetStatusReq) (httpmodels.GetStatusResp, error) {
	var resp httpmodels.GetStatusResp

	jsonResp, err := a.Client.performQuery(http.MethodGet, "/dbfs/get-status", req, nil)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(jsonResp, &resp)
	return resp, err
}

// List lists the contents of a directory, or details of the file
func (a DbfsAPI) List(req httpmodels.ListReq) (httpmodels.ListResp, error) {
	var resp httpmodels.ListResp

	jsonResp, err := a.Client.performQuery(http.MethodGet, "/dbfs/list", req, nil)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(jsonResp, &resp)
	return resp, err
}

// Mkdirs creates the given directory and necessary parent directories if they do not exist
func (a DbfsAPI) Mkdirs(req httpmodels.MkdirsReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/dbfs/mkdirs", req, nil)
	return err
}

// Move moves a file from one location to another location within DBFS
func (a DbfsAPI) Move(req httpmodels.MoveReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/dbfs/move", req, nil)
	return err
}

// Put uploads a file through the use of multipart form post
func (a DbfsAPI) Put(req httpmodels.PutReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/dbfs/put", req, nil)
	return err
}

// Read returns the contents of a file
func (a DbfsAPI) Read(req httpmodels.ReadReq) (httpmodels.ReadResp, error) {
	var resp httpmodels.ReadResp

	jsonResp, err := a.Client.performQuery(http.MethodGet, "/dbfs/read", req, nil)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(jsonResp, &resp)
	return resp, err
}
