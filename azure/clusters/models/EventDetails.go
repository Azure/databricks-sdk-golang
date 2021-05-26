package models

type EventDetails struct {
	CurrentNumWorkers   int32              `json:"current_num_workers,omitempty" url:"current_num_workers,omitempty"`
	TargetNumWorkers    int32              `json:"target_num_workers,omitempty" url:"target_num_workers,omitempty"`
	PreviousAttributes  *ClusterAttributes `json:"previous_attributes,omitempty" url:"previous_attributes,omitempty"`
	Attributes          *ClusterAttributes `json:"attributes,omitempty" url:"attributes,omitempty"`
	PreviousClusterSize ClusterSize        `json:"previous_cluster_size,omitempty" url:"previous_cluster_size,omitempty"`
	ClusterSize         ClusterSize        `json:"cluster_size,omitempty" url:"cluster_size,omitempty"`
}
