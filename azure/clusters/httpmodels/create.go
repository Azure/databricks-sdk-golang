// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package httpmodels

import (
	"github.com/Azure/databricks-sdk-golang/azure/clusters/models"
)

type CreateReq struct {
	NumWorkers             int32                   `json:"num_workers,omitempty" url:"num_workers,omitempty"`
	Autoscale              *models.AutoScale       `json:"autoscale,omitempty" url:"autoscale,omitempty"`
	ClusterName            string                  `json:"cluster_name,omitempty" url:"cluster_name,omitempty"`
	SparkVersion           string                  `json:"spark_version,omitempty" url:"spark_version,omitempty"`
	SparkConf              map[string]string       `json:"spark_conf,omitempty" url:"spark_conf,omitempty"`
	NodeTypeID             string                  `json:"node_type_id,omitempty" url:"node_type_id,omitempty"`
	DriverNodeTypeID       string                  `json:"driver_node_type_id,omitempty" url:"driver_node_type_id,omitempty"`
	CustomTags             []models.ClusterTag     `json:"custom_tags,omitempty" url:"custom_tags,omitempty"`
	ClusterLogConf         *models.ClusterLogConf  `json:"cluster_log_conf,omitempty" url:"cluster_log_conf,omitempty"`
	InitScripts            []models.InitScriptInfo `json:"init_scripts,omitempty" url:"init_scripts,omitempty"`
	DockerImage            *models.DockerImage     `json:"docker_image,omitempty" url:"docker_image,omitempty"`
	SparkEnvVars           map[string]string       `json:"spark_env_vars,omitempty" url:"spark_env_vars,omitempty"`
	AutoterminationMinutes int32                   `json:"autotermination_minutes,omitempty" url:"autotermination_minutes,omitempty"`
	InstancePoolID         string                  `json:"instance_pool_id,omitempty" url:"instance_pool_id,omitempty"`
	IdempotencyToken       string                  `json:"idempotency_token,omitempty" url:"idempotency_token,omitempty"`
	ApplyPolicyDefVal      bool                    `json:"apply_policy_default_values,omitempty" url:"apply_policy_default_values,omitempty"`
	EnableLocalDiskEncr    bool                    `json:"enable_local_disk_encryption,omitempty" url:"enable_local_disk_encryption,omitempty"`
}

type CreateResp struct {
	ClusterID string `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
}
