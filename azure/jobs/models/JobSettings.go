// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package models

import (
	libraryModels "github.com/Azure/databricks-sdk-golang/azure/libraries/models"
)

type JobSettings struct {
	ExistingClusterID      string                   `json:"existing_cluster_id,omitempty" url:"existing_cluster_id,omitempty"`
	NewCluster             *NewCluster              `json:"new_cluster,omitempty" url:"new_cluster,omitempty"`
	NotebookTask           *NotebookTask            `json:"notebook_task,omitempty" url:"notebook_task,omitempty"`
	SparkJarTask           *SparkJarTask            `json:"spark_jar_task,omitempty" url:"spark_jar_task,omitempty"`
	SparkPythonTask        *SparkPythonTask         `json:"spark_python_task,omitempty" url:"spark_python_task,omitempty"`
	SparkSubmitTask        *SparkSubmitTask         `json:"spark_submit_task,omitempty" url:"spark_submit_task,omitempty"`
	Name                   string                   `json:"name,omitempty" url:"name,omitempty"`
	Libraries              *[]libraryModels.Library `json:"libraries,omitempty" url:"libraries,omitempty"`
	EmailNotifications     *JobEmailNotifications   `json:"email_notifications,omitempty" url:"email_notifications,omitempty"`
	TimeoutSeconds         int32                    `json:"timeout_seconds,omitempty" url:"timeout_seconds,omitempty"`
	MaxRetries             int32                    `json:"max_retries,omitempty" url:"max_retries,omitempty"`
	MinRetryIntervalMillis int32                    `json:"min_retry_interval_millis,omitempty" url:"min_retry_interval_millis,omitempty"`
	RetryOnTimeout         bool                     `json:"retry_on_timeout,omitempty" url:"retry_on_timeout,omitempty"`
	Schedule               *CronSchedule            `json:"schedule,omitempty" url:"schedule,omitempty"`
	MaxConcurrentRuns      int32                    `json:"max_concurrent_runs,omitempty" url:"max_concurrent_runs,omitempty"`
}
