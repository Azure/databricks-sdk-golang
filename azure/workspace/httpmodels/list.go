// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package httpmodels

import (
	"github.com/Azure/databricks-sdk-golang/azure/workspace/models"
)

type ListReq struct {
	Path string `json:"path,omitempty" url:"path,omitempty"`
}

type ListResp struct {
	Objects *[]models.ObjectInfo `json:"objects,omitempty" url:"objects,omitempty"`
}
