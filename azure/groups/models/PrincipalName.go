package models

type PrincipalName struct {
	UserName  string `json:"user_name,omitempty" url:"user_name,omitempty"`
	GroupName string `json:"group_name,omitempty" url:"group_name,omitempty"`
}
