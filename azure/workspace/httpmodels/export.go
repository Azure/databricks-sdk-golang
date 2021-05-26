// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package httpmodels

import (
	"github.com/Azure/databricks-sdk-golang/azure/workspace/models"
)

type ExportReq struct {
	Path           string              `json:"path,omitempty" url:"path,omitempty"`
	Format         models.ExportFormat `json:"format,omitempty" url:"format,omitempty"`
	DirectDownload bool                `json:"direct_download,omitempty" url:"direct_download,omitempty"`
}

type ExportResp struct {
	Content string `json:"content,omitempty" url:"content,omitempty"`
}
