package httpmodels

import (
	"github.com/polar-rams/databricks-sdk-golang/azure/clusters/models"
)

type ListNodeTypesResp struct {
	NodeTypes []models.NodeType `json:"node_types,omitempty" url:"node_types,omitempty"`
}
