package models

import (
	clusterModels "github.com/polar-rams/databricks-sdk-golang/azure/clusters/models"
)

type InstancePoolAndStats struct {
	InstancePoolName                   string                      `json:"instance_pool_name,omitempty" url:"instance_pool_name,omitempty"`
	MinIdleInstances                   int32                       `json:"min_idle_instances,omitempty" url:"min_idle_instances,omitempty"`
	MaxCapacity                        int32                       `json:"max_capacity,omitempty" url:"max_capacity,omitempty"`
	NodeTypeID                         string                      `json:"node_type_id,omitempty" url:"node_type_id,omitempty"`
	CustomTags                         *[]clusterModels.ClusterTag `json:"custom_tags,omitempty" url:"custom_tags,omitempty"`
	IdleInstanceAutoterminationMinutes int32                       `json:"idle_instance_autotermination_minutes,omitempty" url:"idle_instance_autotermination_minutes,omitempty"`
	EnableElasticDisk                  bool                        `json:"enable_elastic_disk,omitempty" url:"enable_elastic_disk,omitempty"`
	DiskSpec                           DiskSpec                    `json:"disk_spec,omitempty" url:"disk_spec,omitempty"`
	PreloadedSparkVersions             *[]string                   `json:"preloaded_spark_versions,omitempty" url:"preloaded_spark_versions,omitempty"`
	InstancePoolID                     string                      `json:"instance_pool_id,omitempty" url:"instance_pool_id,omitempty"`
	DefaultTags                        *[]clusterModels.ClusterTag `json:"default_tags,omitempty" url:"default_tags,omitempty"`
	State                              InstancePoolState           `json:"state,omitempty" url:"state,omitempty"`
	Stats                              InstancePoolStats           `json:"stats,omitempty" url:"stats,omitempty"`
}
