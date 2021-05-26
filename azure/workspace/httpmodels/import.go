// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package httpmodels

import (
	"github.com/Azure/databricks-sdk-golang/azure/workspace/models"
)

type ImportReq struct {
	Path      string              `json:"path,omitempty" url:"path,omitempty"`
	Format    models.ExportFormat `json:"format,omitempty" url:"format,omitempty"`
	Language  models.Language     `json:"language,omitempty" url:"language,omitempty"`
	Content   string              `json:"content,omitempty" url:"content,omitempty"`
	Overwrite bool                `json:"overwrite,omitempty" url:"overwrite,omitempty"`
}
