// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package httpmodels

import (
	"github.com/Azure/databricks-sdk-golang/azure/dbfs/models"
)

type ListReq struct {
	Path string `json:"path,omitempty" url:"path,omitempty"`
}

type ListResp struct {
	Files *[]models.FileInfo `json:"files,omitempty" url:"files,omitempty"`
}
