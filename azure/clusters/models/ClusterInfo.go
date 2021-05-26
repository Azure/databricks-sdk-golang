// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package models

type ClusterInfo struct {
	NumWorkers             int32             `json:"num_workers,omitempty" url:"num_workers,omitempty"`
	AutoScale              *AutoScale        `json:"autoscale,omitempty" url:"autoscale,omitempty"`
	ClusterID              string            `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
	CreatorUserName        string            `json:"creator_user_name,omitempty" url:"creator_user_name,omitempty"`
	Driver                 SparkNode         `json:"driver,omitempty" url:"driver,omitempty"`
	Executors              *[]SparkNode      `json:"executors,omitempty" url:"executors,omitempty"`
	SparkContextID         int64             `json:"spark_context_id,omitempty" url:"spark_context_id,omitempty"`
	JdbcPort               int32             `json:"jdbc_port,omitempty" url:"jdbc_port,omitempty"`
	ClusterName            string            `json:"cluster_name,omitempty" url:"cluster_name,omitempty"`
	SparkVersion           string            `json:"spark_version,omitempty" url:"spark_version,omitempty"`
	SparkConf              SparkConfPair     `json:"spark_conf,omitempty" url:"spark_conf,omitempty"`
	NodeTypeID             string            `json:"node_type_id,omitempty" url:"node_type_id,omitempty"`
	DriverNodeTypeID       string            `json:"driver_node_type_id,omitempty" url:"driver_node_type_id,omitempty"`
	ClusterLogConf         ClusterLogConf    `json:"cluster_log_conf,omitempty" url:"cluster_log_conf,omitempty"`
	InitScripts            *[]InitScriptInfo `json:"init_scripts,omitempty" url:"init_scripts,omitempty"`
	SparkEnvVars           map[string]string `json:"spark_env_vars,omitempty" url:"spark_env_vars,omitempty"`
	AutoterminationMinutes int32             `json:"autotermination_minutes,omitempty" url:"autotermination_minutes,omitempty"`
	State                  ClusterState      `json:"state,omitempty" url:"state,omitempty"`
	StateMessage           string            `json:"state_message,omitempty" url:"state_message,omitempty"`
	StartTime              int64             `json:"start_time,omitempty" url:"start_time,omitempty"`
	TerminateTime          int64             `json:"terminate_time,omitempty" url:"terminate_time,omitempty"`
	LastStateLossTime      int64             `json:"last_state_loss_time,omitempty" url:"last_state_loss_time,omitempty"`
	LastActivityTime       int64             `json:"last_activity_time,omitempty" url:"last_activity_time,omitempty"`
	ClusterMemoryMb        int64             `json:"cluster_memory_mb,omitempty" url:"cluster_memory_mb,omitempty"`
	ClusterCores           float32           `json:"cluster_cores,omitempty" url:"cluster_cores,omitempty"`
	DefaultTags            map[string]string `json:"default_tags,omitempty" url:"default_tags,omitempty"`
	ClusterLogStatus       LogSyncStatus     `json:"cluster_log_status,omitempty" url:"cluster_log_status,omitempty"`
	TerminationReason      TerminationReason `json:"termination_reason,omitempty" url:"termination_reason,omitempty"`
}
