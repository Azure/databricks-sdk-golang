// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package httpmodels

type DeleteSecretACLReq struct {
	Scope     string `json:"scope,omitempty" url:"scope,omitempty"`
	Principal string `json:"principal,omitempty" url:"principal,omitempty"`
}
