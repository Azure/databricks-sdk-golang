package httpmodels

import "github.com/polar-rams/databricks-sdk-golang/azure/secrets/models"

type ListSecretScopesResp struct {
	Scopes *[]models.SecretScope `json:"scopes,omitempty" url:"scopes,omitempty"`
}
