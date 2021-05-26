// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package httpmodels

import "github.com/Azure/databricks-sdk-golang/azure/jobs/models"

type GetReq struct {
	JobID int64 `json:"job_id,omitempty" url:"job_id,omitempty"`
}

type GetResp struct {
	JobID            int64               `json:"job_id,omitempty" url:"job_id,omitempty"`
	CreateorUserName string              `json:"creator_user_name,omitempty" url:"creator_user_name,omitempty"`
	Settings         *models.JobSettings `json:"settings,omitempty" url:"settings,omitempty"`
	CreatedTime      int64               `json:"created_time,omitempty" url:"created_time,omitempty"`
}
