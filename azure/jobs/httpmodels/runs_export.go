// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package httpmodels

import "github.com/Azure/databricks-sdk-golang/azure/jobs/models"

type RunsExportReq struct {
	RunID         int64                `json:"run_id,omitempty" url:"run_id,omitempty"`
	ViewsToExport models.ViewsToExport `json:"views_to_export,omitempty" url:"views_to_export,omitempty"`
}

type RunsExportResp struct {
	Views *[]models.ViewItem `json:"views,omitempty" url:"views,omitempty"`
}
