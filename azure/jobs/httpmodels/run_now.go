// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package httpmodels

type RunNowReq struct {
	JobID             int64              `json:"job_id,omitempty" url:"job_id,omitempty"`
	JarParams         *[]string          `json:"jar_params,omitempty" url:"jar_params,omitempty"`
	NotebookParams    *map[string]string `json:"notebook_params,omitempty" url:"notebook_params,omitempty"`
	PythonParams      *[]string          `json:"python_params,omitempty" url:"python_params,omitempty"`
	SparkSubmitParams *[]string          `json:"spark_submit_params,omitempty" url:"spark_submit_params,omitempty"`
}

type RunNowResp struct {
	RunID       int64 `json:"run_id,omitempty" url:"run_id,omitempty"`
	NumberInJob int64 `json:"number_in_job,omitempty" url:"number_in_job,omitempty"`
}
