// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package httpmodels

import (
	"github.com/Azure/databricks-sdk-golang/azure/workspace/models"
)

type GetStatusReq struct {
	Path string `json:"path,omitempty" url:"path,omitempty"`
}

type GetStatusResp struct {
	ObjectType models.ObjectType `json:"object_type,omitempty" url:"object_type,omitempty"`
	ObjectID   int64             `json:"object_id,omitempty" url:"object_id,omitempty"`
	Path       string            `json:"path,omitempty" url:"path,omitempty"`
	Language   models.Language   `json:"language,omitempty" url:"language,omitempty"`
}
