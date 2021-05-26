// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package httpmodels

import "github.com/Azure/databricks-sdk-golang/azure/libraries/models"

type AllClusterStatusesResp struct {
	Statuses *[]models.ClusterLibraryStatuses `json:"statuses,omitempty" url:"statuses,omitempty"`
}
