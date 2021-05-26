package models

type ClusterInstance struct {
	ClusterID      string `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
	SparkContextID string `json:"spark_context_id,omitempty" url:"spark_context_id,omitempty"`
}
