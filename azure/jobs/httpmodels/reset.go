// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package httpmodels

import "github.com/Azure/databricks-sdk-golang/azure/jobs/models"

type ResetReq struct {
	JobID       int64               `json:"job_id,omitempty" url:"job_id,omitempty"`
	NewSettings *models.JobSettings `json:"new_settings,omitempty" url:"new_settings,omitempty"`
}
