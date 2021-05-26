// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package httpmodels

type AddBlockReq struct {
	Handle int64  `json:"handle,omitempty" url:"handle,omitempty"`
	Data   string `json:"data,omitempty" url:"data,omitempty"`
}