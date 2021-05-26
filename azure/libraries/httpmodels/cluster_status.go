package httpmodels

import "github.com/polar-rams/databricks-sdk-golang/azure/libraries/models"

type ClusterStatusReq struct {
	ClusterID string `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
}

type ClusterStatusResp struct {
	ClusterID      string                      `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
	LibraryStatues *[]models.LibraryFullStatus `json:"library_statuses,omitempty" url:"library_statuses,omitempty"`
}
