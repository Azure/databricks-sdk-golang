// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package httpmodels

import "github.com/Azure/databricks-sdk-golang/azure/clusters/models"

type RuntimeVersionsResp struct {
	Versions []models.SparkVersion `json:"versions,omitempty" url:"versions,omitempty"`
}
