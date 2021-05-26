package httpmodels

import "github.com/polar-rams/databricks-sdk-golang/azure/libraries/models"

type AllClusterStatusesResp struct {
	Statuses *[]models.ClusterLibraryStatuses `json:"statuses,omitempty" url:"statuses,omitempty"`
}
