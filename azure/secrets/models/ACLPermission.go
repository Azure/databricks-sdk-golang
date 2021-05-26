package models

type ACLPermission string

const (
	ACLPermissionRead   = "READ"
	ACLPermissionWrite  = "WRITE"
	ACLPermissionManage = "MANAGE"
)
