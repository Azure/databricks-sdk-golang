package httpmodels

import (
	"github.com/polar-rams/databricks-sdk-golang/azure/clusters/models"
)

type ResizeReq struct {
	ClusterID  string           `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
	NumWorkers int32            `json:"num_workers,omitempty" url:"num_workers,omitempty"`
	Autoscale  models.AutoScale `json:"autoscale,omitempty" url:"autoscale,omitempty"`
}
