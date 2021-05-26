// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package httpmodels

import "github.com/Azure/databricks-sdk-golang/azure/secrets/models"

type ListSecretScopesResp struct {
	Scopes *[]models.SecretScope `json:"scopes,omitempty" url:"scopes,omitempty"`
}
