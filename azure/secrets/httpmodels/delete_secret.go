// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package httpmodels

type DeleteSecretReq struct {
	Scope string `json:"scope,omitempty" url:"scope,omitempty"`
	Key   string `json:"key,omitempty" url:"key,omitempty"`
}
