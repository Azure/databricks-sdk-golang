// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package httpmodels

import "github.com/Azure/databricks-sdk-golang/azure/secrets/models"

type GetSecretACLReq struct {
	Scope     string `json:"scope,omitempty" url:"scope,omitempty"`
	Principal string `json:"principal,omitempty" url:"principal,omitempty"`
}

type GetSecretACLResp struct {
	Principal  string               `json:"principal,omitempty" url:"principal,omitempty"`
	Permission models.ACLPermission `json:"permission,omitempty" url:"permission,omitempty"`
}
