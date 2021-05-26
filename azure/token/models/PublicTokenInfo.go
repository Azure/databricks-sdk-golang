// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package models

type PublicTokenInfo struct {
	TokenID      string `json:"token_id,omitempty" url:"token_id,omitempty"`
	CreationTime int64  `json:"creation_time,omitempty" url:"creation_time,omitempty"`
	ExpiryTime   int64  `json:"expiry_time,omitempty" url:"expiry_time,omitempty"`
	Comment      string `json:"comment,omitempty" url:"comment,omitempty"`
}
