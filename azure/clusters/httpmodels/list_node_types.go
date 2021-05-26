// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package httpmodels

import (
	"github.com/Azure/databricks-sdk-golang/azure/clusters/models"
)

type ListNodeTypesResp struct {
	NodeTypes []models.NodeType `json:"node_types,omitempty" url:"node_types,omitempty"`
}
