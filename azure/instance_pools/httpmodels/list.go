// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package httpmodels

import (
	"github.com/Azure/databricks-sdk-golang/azure/instance_pools/models"
)

//ListResp is to compose the response to list an array of instance pools' stats
type ListResp struct {
	InstancePools *[]models.InstancePoolStats `json:"instance_pools,omitempty" url:"instance_pools,omitempty"`
}
