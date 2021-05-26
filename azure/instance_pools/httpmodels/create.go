package httpmodels

import (
	clustersModels "github.com/polar-rams/databricks-sdk-golang/azure/clusters/models"
	"github.com/polar-rams/databricks-sdk-golang/azure/instance_pools/models"
)

//CreateReq is to compose the request for creating an instance pool
type CreateReq struct {
	InstancePoolName                   string                             `json:"instance_pool_name,omitempty" url:"instance_pool_name,omitempty"`
	MinIdleInstances                   int32                              `json:"min_idle_instances,omitempty" url:"min_idle_instances,omitempty"`
	MaxCapacity                        int32                              `json:"max_capacity,omitempty" url:"max_capacity,omitempty"`
	NodeTypeID                         string                             `json:"node_type_id,omitempty" url:"node_type_id,omitempty"`
	CustomTags                         *[]clustersModels.ClusterTag       `json:"custom_tags,omitempty" url:"custom_tags,omitempty"`
	IdleInstanceAutoterminationMinutes int32                              `json:"idle_instance_autotermination_minutes,omitempty" url:"idle_instance_autotermination_minutes,omitempty"`
	EnableElasticDisk                  bool                               `json:"enable_elastic_disk,omitempty" url:"enable_elastic_disk,omitempty"`
	DiskSpec                           models.DiskSpec                    `json:"disk_spec,omitempty" url:"disk_spec,omitempty"`
	PreloadedSparkVersions             *[]string                          `json:"preloaded_spark_versions,omitempty" url:"preloaded_spark_versions,omitempty"`
	PreloadedDockerImages              *[]clustersModels.DockerImage      `json:"preloaded_docker_images,omitempty" url:"preloaded_docker_images,omitempty"`
	AzureAttributes                    models.InstancePoolAzureAttributes `json:"azure_attributes,omitempty" url:"azure_attributes,omitempty"`
}

//CreateResp to compose the response of creation request
type CreateResp struct {
	InstancePoolID string `json:"instance_pool_id,omitempty" url:"instance_pool_id,omitempty"`
}
