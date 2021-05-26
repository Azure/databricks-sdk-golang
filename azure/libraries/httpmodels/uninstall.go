// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package httpmodels

import "github.com/Azure/databricks-sdk-golang/azure/libraries/models"

type UninstallReq struct {
	ClusterID string            `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
	Libraries *[]models.Library `json:"libraries,omitempty" url:"libraries,omitempty"`
}
