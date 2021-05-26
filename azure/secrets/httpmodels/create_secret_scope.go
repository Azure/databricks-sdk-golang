package httpmodels

import "github.com/polar-rams/databricks-sdk-golang/azure/secrets/models"

type CreateSecretScopeReq struct {
	Scope                string                  `json:"scope,omitempty" url:"scope,omitempty"`
	ScopeBackendType     models.ScopeBackendType `json:"scope_backend_type,omitempty" url:"scope_backend_type,omitempty"`
	BackendAzureKeyvault struct {
		ResourceID string `json:"backend_azure_keyvault,omitempty" url:"backend_azure_keyvault,omitempty"`
		DNSName    string `json:"dns_name,omitempty" url:"dns_name,omitempty"`
	}
	InitialManagePrincipal string `json:"initial_manage_principal,omitempty" url:"initial_manage_principal,omitempty"`
}
