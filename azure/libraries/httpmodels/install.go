package httpmodels

import "github.com/polar-rams/databricks-sdk-golang/azure/libraries/models"

type InstallReq struct {
	ClusterID string            `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
	Libraries *[]models.Library `json:"libraries,omitempty" url:"libraries,omitempty"`
}
