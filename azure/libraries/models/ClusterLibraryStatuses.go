package models

type ClusterLibraryStatuses struct {
	ClusterID       string               `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
	LibraryStatuses *[]LibraryFullStatus `json:"library_statuses,omitempty" url:"library_statuses,omitempty"`
}
