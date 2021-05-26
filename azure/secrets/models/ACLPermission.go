// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package models

type ACLPermission string

const (
	ACLPermissionRead   = "READ"
	ACLPermissionWrite  = "WRITE"
	ACLPermissionManage = "MANAGE"
)
