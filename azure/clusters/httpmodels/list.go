// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package httpmodels

import "github.com/Azure/databricks-sdk-golang/azure/clusters/models"

type ListResp struct {
	Clusters []models.ClusterInfo
}
