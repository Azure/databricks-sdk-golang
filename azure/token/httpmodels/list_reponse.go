// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package httpmodels

import (
	"github.com/Azure/databricks-sdk-golang/azure/token/models"
)

type ListResp struct {
	TokenInfos *[]models.PublicTokenInfo `json:"token_infos,omitempty" url:"token_infos,omitempty"`
}
