// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package httpmodels

type RemoveMemberReq struct {
	UserName   string `json:"user_name,omitempty" url:"user_name,omitempty"`
	GroupName  string `json:"group_name,omitempty" url:"group_name,omitempty"`
	ParentName string `json:"parent_name,omitempty" url:"parent_name,omitempty"`
}
