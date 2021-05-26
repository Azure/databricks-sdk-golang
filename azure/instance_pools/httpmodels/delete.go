// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package httpmodels

// DeleteReq is to compose the request to delete an instance pool
type DeleteReq struct {
	InstancePoolID string `json:"instance_pool_id,omitempty" url:"instance_pool_id,omitempty"`
}
