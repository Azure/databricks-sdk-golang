package httpmodels

import "github.com/polar-rams/databricks-sdk-golang/azure/secrets/models"

type PutSecretACLReq struct {
	Scope      string               `json:"scope,omitempty" url:"scope,omitempty"`
	Principal  string               `json:"principal,omitempty" url:"principal,omitempty"`
	Permission models.ACLPermission `json:"permission,omitempty" url:"permission,omitempty"`
}
