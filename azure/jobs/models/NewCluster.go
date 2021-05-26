// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package models

import (
	clusterModels "github.com/Azure/databricks-sdk-golang/azure/clusters/models"
)

type NewCluster struct {
	NumWorkers        int32                           `json:"num_workers,omitempty" url:"num_workers,omitempty"`
	Autoscale         *clusterModels.AutoScale        `json:"autoscale,omitempty" url:"autoscale,omitempty"`
	SparkVersion      string                          `json:"spark_version,omitempty" url:"spark_version,omitempty"`
	SparkConf         *map[string]string              `json:"spark_conf,omitempty" url:"spark_conf,omitempty"`
	NodeTypeID        string                          `json:"node_type_id,omitempty" url:"node_type_id,omitempty"`
	DriverNodeTypeID  string                          `json:"driver_node_type_id,omitempty" url:"driver_node_type_id,omitempty"`
	CustomTags        *[]clusterModels.ClusterTag     `json:"custom_tags,omitempty" url:"custom_tags,omitempty"`
	ClusterLogConf    *clusterModels.ClusterLogConf   `json:"cluster_log_conf,omitempty" url:"cluster_log_conf,omitempty"`
	InitScripts       *[]clusterModels.InitScriptInfo `json:"init_scripts,omitempty" url:"init_scripts,omitempty"`
	SparkEnvVars      *map[string]string              `json:"spark_env_vars,omitempty" url:"spark_env_vars,omitempty"`
	EnableElasticDisk bool                            `json:"enable_elastic_disk,omitempty" url:"enable_elastic_disk,omitempty"`
	InstancePoolID    string                          `json:"instance_pool_id,omitempty" url:"instance_pool_id,omitempty"`
}
