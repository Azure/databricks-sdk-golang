// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package httpmodels

import (
	"github.com/Azure/databricks-sdk-golang/azure/jobs/models"
)

type RunsGetOutputReq struct {
	RunID int64 `json:"run_id,omitempty" url:"run_id,omitempty"`
}

type RunsGetOutputResp struct {
	NotebookOutput models.NotebookOutput `json:"notebook_output,omitempty" url:"notebook_output,omitempty"`
	Error          string                `json:"error,omitempty" url:"error,omitempty"`
	Metadata       models.Run            `json:"metadata,omitempty" url:"metadata,omitempty"`
}
