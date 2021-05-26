// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package httpmodels

import (
	"github.com/Azure/databricks-sdk-golang/azure/jobs/models"
)

type ListResp struct {
	Jobs *[]models.Job `json:"jobs,omitempty" url:"jobs,omitempty"`
}
