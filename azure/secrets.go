package azure

import (
	"encoding/json"
	"net/http"

	"github.com/polar-rams/databricks-sdk-golang/azure/secrets/httpmodels"
)

// SecretsAPI exposes the Secrets API
type SecretsAPI struct {
	Client DBClient
}

func (a SecretsAPI) init(client DBClient) SecretsAPI {
	a.Client = client
	return a
}

// CreateSecretScope create an Azure Key Vault-backed or Databricks-backed scope
func (a SecretsAPI) CreateSecretScope(req httpmodels.CreateSecretScopeReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/secrets/scopes/create", req, nil)
	return err
}

// DeleteSecretScope deletes a secret scope
func (a SecretsAPI) DeleteSecretScope(req httpmodels.DeleteSecretScopeReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/secrets/scopes/delete", req, nil)
	return err
}

// ListSecretScopes lists all secret scopes available in the workspace
func (a SecretsAPI) ListSecretScopes() (httpmodels.ListSecretScopesResp, error) {
	var resp httpmodels.ListSecretScopesResp

	jsonResp, err := a.Client.performQuery(http.MethodGet, "/secrets/scopes/list", nil, nil)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(jsonResp, &resp)
	return resp, err
}

// PutSecret creates or modifies a bytes secret depends on the type of scope backend with
func (a SecretsAPI) PutSecret(req httpmodels.PutSecretReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/secrets/put", req, nil)
	return err
}

// DeleteSecret deletes a secret depends on the type of scope backend
func (a SecretsAPI) DeleteSecret(req httpmodels.DeleteSecretReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/secrets/delete", req, nil)
	return err
}

// ListSecrets lists the secret keys that are stored at this scope
func (a SecretsAPI) ListSecrets(req httpmodels.ListSecretsReq) (httpmodels.ListSecretsResp, error) {
	var resp httpmodels.ListSecretsResp

	jsonResp, err := a.Client.performQuery(http.MethodGet, "/secrets/list", req, nil)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(jsonResp, &resp)
	return resp, err
}

// PutSecretACL creates or overwrites the ACL associated with the given principal (user or group) on the specified scope point
func (a SecretsAPI) PutSecretACL(req httpmodels.PutSecretACLReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/secrets/acls/put", req, nil)
	return err
}

// DeleteSecretACL deletes the given ACL on the given scope
func (a SecretsAPI) DeleteSecretACL(req httpmodels.DeleteSecretACLReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/secrets/acls/delete", req, nil)
	return err
}

// GetSecretACL describe the details about the given ACL, such as the group and permission
func (a SecretsAPI) GetSecretACL(req httpmodels.GetSecretACLReq) (httpmodels.GetSecretACLResp, error) {
	var resp httpmodels.GetSecretACLResp

	jsonResp, err := a.Client.performQuery(http.MethodGet, "/secrets/acls/get", req, nil)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(jsonResp, &resp)
	return resp, err
}

// ListSecretACLs lists the ACLs set on the given scope
func (a SecretsAPI) ListSecretACLs(req httpmodels.ListSecretACLsReq) (httpmodels.ListSecretACLsResp, error) {
	var resp httpmodels.ListSecretACLsResp

	jsonResp, err := a.Client.performQuery(http.MethodGet, "/secrets/acls/list", req, nil)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(jsonResp, &resp)
	return resp, err
}
