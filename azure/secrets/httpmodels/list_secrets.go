package httpmodels

import "github.com/polar-rams/databricks-sdk-golang/azure/secrets/models"

type ListSecretsReq struct {
	Scope string `json:"scope,omitempty" url:"scope,omitempty"`
}

type ListSecretsResp struct {
	Secrets []models.SecretMetadata `json:"secrets,omitempty" url:"secrets,omitempty"`
}
