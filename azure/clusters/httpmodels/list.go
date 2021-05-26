package httpmodels

import "github.com/polar-rams/databricks-sdk-golang/azure/clusters/models"

type ListResp struct {
	Clusters []models.ClusterInfo
}
