// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package httpmodels

import (
	"github.com/Azure/databricks-sdk-golang/azure/clusters/models"
)

type EditReq struct {
	NumWorkers             int32                   `json:"num_workers,omitempty" url:"num_workers,omitempty"`
	AutoScale              models.AutoScale        `json:"autoscale,omitempty" url:"autoscale,omitempty"`
	ClusterID              string                  `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
	ClusterName            string                  `json:"cluster_name,omitempty" url:"cluster_name,omitempty"`
	SparkVersion           string                  `json:"spark_version,omitempty" url:"spark_version,omitempty"`
	SparkConf              models.SparkConfPair    `json:"spark_conf,omitempty" url:"spark_conf,omitempty"`
	DockerImage            models.DockerImage      `json:"docker_image,omitempty" url:"docker_image,omitempty"`
	NodeTypeID             string                  `json:"node_type_id,omitempty" url:"node_type_id,omitempty"`
	InstancePoolID         string                  `json:"instance_pool_id,omitempty" url:"instance_pool_id,omitempty"`
	DriverNodeTypeID       string                  `json:"driver_node_type_id,omitempty" url:"driver_node_type_id,omitempty"`
	ClusterLogConf         models.ClusterLogConf   `json:"cluster_log_conf,omitempty" url:"cluster_log_conf,omitempty"`
	InitScripts            []models.InitScriptInfo `json:"init_scripts,omitempty" url:"init_scripts,omitempty"`
	SparkEnvVars           map[string]string       `json:"spark_env_vars,omitempty" url:"spark_env_vars,omitempty"`
	AutoterminationMinutes int32                   `json:"autotermination_minutes,omitempty" url:"autotermination_minutes,omitempty"`
	ApplyPolicyDefVal      bool                    `json:"apply_policy_default_values,omitempty" url:"apply_policy_default_values,omitempty"`
	EnableLocalDiskEncr    bool                    `json:"enable_local_disk_encryption,omitempty" url:"enable_local_disk_encryption,omitempty"`
}
