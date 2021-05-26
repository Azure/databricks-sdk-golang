// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package httpmodels

import (
	"github.com/Azure/databricks-sdk-golang/azure/jobs/models"
	libraryModels "github.com/Azure/databricks-sdk-golang/azure/libraries/models"
)

type RunsSubmitReq struct {
	ExistingCluster  string                   `json:"run_id,omitempty" url:"run_id,omitempty"`
	NewCluster       *models.NewCluster       `json:"new_cluster,omitempty" url:"new_cluster,omitempty"`
	NotebookTask     *models.NotebookTask     `json:"notebook_task,omitempty" url:"notebook_task,omitempty"`
	SparkJarTask     *models.SparkJarTask     `json:"spark_jar_task,omitempty" url:"spark_jar_task,omitempty"`
	SparkPythonTask  *models.SparkPythonTask  `json:"spark_python_task,omitempty" url:"spark_python_task,omitempty"`
	SparkSubmitTask  *models.SparkSubmitTask  `json:"spark_submit_task,omitempty" url:"spark_submit_task,omitempty"`
	RunName          string                   `json:"run_name,omitempty" url:"run_name,omitempty"`
	Libraries        *[]libraryModels.Library `json:"libraries,omitempty" url:"libraries,omitempty"`
	TimeoutSeconds   int32                    `json:"timeout_seconds,omitempty" url:"timeout_seconds,omitempty"`
	IdempotencyToken string                   `json:"idempotency_token,omitempty" url:"idempotency_token,omitempty"`
}

type RunsSubmitResp struct {
	RunID int64 `json:"run_id,omitempty" url:"run_id,omitempty"`
}
