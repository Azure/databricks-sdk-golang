package models

type SecretScope struct {
	Name        string           `json:"name,omitempty" url:"name,omitempty"`
	BackendType ScopeBackendType `json:"backend_type,omitempty" url:"backend_type,omitempty"`
}
