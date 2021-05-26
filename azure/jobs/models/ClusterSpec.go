package models

import (
	libraryModels "github.com/polar-rams/databricks-sdk-golang/azure/libraries/models"
)

type ClusterSpec struct {
	ExistingClusterID string                   `json:"existing_cluster_id,omitempty" url:"existing_cluster_id,omitempty"`
	NewCluster        NewCluster               `json:"new_cluster,omitempty" url:"new_cluster,omitempty"`
	Libraries         *[]libraryModels.Library `json:"libraries,omitempty" url:"libraries,omitempty"`
}
