// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package models

type RunParameters struct {
	JarParams         *[]string          `json:"jar_params,omitempty" url:"jar_params,omitempty"`
	NotebookParams    *map[string]string `json:"notebook_params,omitempty" url:"notebook_params,omitempty"`
	PythonParams      *[]string          `json:"python_params,omitempty" url:"python_params,omitempty"`
	SparkSubmitParams *[]string          `json:"spark_submit_params,omitempty" url:"spark_submit_params,omitempty"`
}
