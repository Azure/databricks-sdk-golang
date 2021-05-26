// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package httpmodels

import "github.com/Azure/databricks-sdk-golang/azure/groups/models"

type ListMembersReq struct {
	GroupName string `json:"group_name,omitempty" url:"group_name,omitempty"`
}

type ListMembersResp struct {
	Members *[]models.PrincipalName `json:"members,omitempty" url:"members,omitempty"`
}
