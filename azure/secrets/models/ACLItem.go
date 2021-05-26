package models

type ACLItem struct {
	Principal  string        `json:"principal,omitempty" url:"principal,omitempty"`
	Permission ACLPermission `json:"permission,omitempty" url:"permission,omitempty"`
}
