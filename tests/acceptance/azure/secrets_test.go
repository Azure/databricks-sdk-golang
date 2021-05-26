package azure_test

import (
	"fmt"
	"testing"

	"github.com/polar-rams/databricks-sdk-golang/azure/secrets/httpmodels"
	"github.com/polar-rams/databricks-sdk-golang/azure/secrets/models"
	"github.com/stretchr/testify/assert"
)

const (
	scpNmPfx = "testSecret"
	scpUsers = "users"
)

func TestAzureSecretsScopeCreateListDelete(t *testing.T) {

	scpNm := fmt.Sprintf("%s%s%s", scpNmPfx, "SS", randSeq(6))

	// Test Create Secret Scope and assert
	csr := httpmodels.CreateSecretScopeReq{
		Scope:                  scpNm,
		InitialManagePrincipal: scpUsers,
	}
	e := c.Secrets().CreateSecretScope(csr)
	assert.Nil(t, e, fmt.Sprintf("could not create secret scope %s", scpNm))

	// Test List Secret Scopes and assert scope exists
	var lscr httpmodels.ListSecretScopesResp
	lscr, e = c.Secrets().ListSecretScopes()
	assert.Nil(t, e, "could not list scopes")

	sc := lscr.Scopes
	assert.Equal(t, SctScpExists(*sc, scpNm), true, fmt.Sprintf("expected secret scope %s to exist", scpNm))

	// Test Delete Secret Scope
	dscr := httpmodels.DeleteSecretScopeReq{
		Scope: scpNm,
	}
	e = c.Secrets().DeleteSecretScope(dscr)
	assert.Nil(t, e, fmt.Sprintf("could not delete secret scope : %s", scpNm))

	lscr, e = c.Secrets().ListSecretScopes()
	assert.Nil(t, e, "could not list scopes after deletion")

	if sc = lscr.Scopes; sc != nil {
		assert.NotEqual(t, SctScpExists(*sc, scpNm), true, fmt.Sprintf("expected secret scope %s to not exist after deletion", scpNm))
	}

}

func TestAzureSecretsPutListDelete(t *testing.T) {

	// setup random suffix for test object names
	nmSfx := randSeq(6)

	scpNm := fmt.Sprintf("%s%s%s", scpNmPfx, "SA", nmSfx)

	// Setup Secret Scope for secrets and assert
	csr := httpmodels.CreateSecretScopeReq{
		Scope:                  scpNm,
		InitialManagePrincipal: scpUsers,
	}
	e := c.Secrets().CreateSecretScope(csr)
	assert.Nil(t, e, fmt.Sprintf("could not create secret scope %s", scpNm))

	// Put Secret and Assert
	sctKy := fmt.Sprintf("%s%s", "sec", nmSfx)
	psr := httpmodels.PutSecretReq{
		Scope:       scpNm,
		Key:         sctKy,
		StringValue: fmt.Sprintf("%s%s", nmSfx, nmSfx),
	}
	e = c.Secrets().PutSecret(psr)
	assert.Nil(t, e, fmt.Sprintf("could not put secret %s (%s)", sctKy, scpNm))

	// List Secret and Assert
	lsr := httpmodels.ListSecretsReq{
		Scope: scpNm,
	}
	lsrsp, e := c.Secrets().ListSecrets(lsr)
	assert.Nil(t, e, fmt.Sprintf("could not list secrets in scope %s", scpNm))

	s := lsrsp.Secrets
	assert.Equal(t, SctExists(s, sctKy), true, fmt.Sprintf("could not list secret with Key %s", sctKy))

	// Test Delete Secret
	dsr := httpmodels.DeleteSecretReq{
		Scope: scpNm,
		Key:   sctKy,
	}
	e = c.Secrets().DeleteSecret(dsr)
	assert.Nil(t, e, fmt.Sprintf("could not delete secret  : %s", sctKy))

	lsr = httpmodels.ListSecretsReq{
		Scope: scpNm,
	}
	lsrsp, e = c.Secrets().ListSecrets(lsr)
	assert.Nil(t, e, fmt.Sprintf("could not list secrets after secret deletion in scope %s", scpNm))

	if s = lsrsp.Secrets; s != nil {
		assert.NotEqual(t, SctExists(s, sctKy), true, fmt.Sprintf("Expected secret %s to not exist after deletion", sctKy))
	}
}

func TestAzureSecretsACLPutGetListDelete(t *testing.T) {

	const (
		aclPrin = "admins"
		aclPerm = "READ"
	)

	// setup random suffix for test object names
	nmSfx := randSeq(6)

	scpNm := fmt.Sprintf("%s%s%s", scpNmPfx, "SE", nmSfx)

	// Setup Secret Scope for secrets and assert
	csr := httpmodels.CreateSecretScopeReq{
		Scope:                  scpNm,
		InitialManagePrincipal: scpUsers,
	}
	e := c.Secrets().CreateSecretScope(csr)
	assert.Nil(t, e, fmt.Sprintf("could not create secret scope %s", scpNm))

	// Put Secret ACL and assert
	psr := httpmodels.PutSecretACLReq{
		Scope:      scpNm,
		Principal:  aclPrin,
		Permission: aclPerm,
	}
	e = c.Secrets().PutSecretACL(psr)
	assert.Nil(t, e, fmt.Sprintf("could not put secret acl  for scope: %s, Principal: %s, Permission: %s", scpNm, aclPrin, aclPerm))

	// Get Secret ACL and assert
	gsr := httpmodels.GetSecretACLReq{
		Scope:     scpNm,
		Principal: aclPrin,
	}

	gsrsp, e := c.Secrets().GetSecretACL(gsr)
	assert.Nil(t, e, "could not get secret ACL")
	assert.Equal(t, string(gsrsp.Permission), aclPerm, "secret ACL permission not as expected")

	// List Secret ACL and Assert
	lsr := httpmodels.ListSecretACLsReq{
		Scope: scpNm,
	}
	lsrsp, e := c.Secrets().ListSecretACLs(lsr)
	assert.Nil(t, e, fmt.Sprintf("could not list secrets acls in scope %s", scpNm))

	s := lsrsp.Items
	assert.Equal(t, GetSctACLPermission(s, aclPrin), "READ", "ACL permission not as Expected")

	// Test Delete Secret ACL and Assert
	dsr := httpmodels.DeleteSecretACLReq{
		Scope:     scpNm,
		Principal: aclPrin,
	}
	e = c.Secrets().DeleteSecretACL(dsr)
	assert.Nil(t, e, fmt.Sprintf("could not delete secret acl for scope: %s, Principal: %s", scpNm, aclPrin))

	lsr = httpmodels.ListSecretACLsReq{
		Scope: scpNm,
	}
	lsrsp, e = c.Secrets().ListSecretACLs(lsr)
	assert.Nil(t, e, fmt.Sprintf("could not list secrets acls in scope %s after deletion", scpNm))

	s = lsrsp.Items
	if s != nil {
		assert.NotEqual(t, GetSctACLPermission(s, aclPrin), "READ", "Did not expect ACL permission to exist after deletion ")
	}
}

// SctScpExists checks if scope name exists in secret scopes
func SctScpExists(sc []models.SecretScope, scpNm string) bool {
	for _, scp := range sc {
		if scp.Name == scpNm {
			return true
		}
	}
	return false
}

// SctExists checks if secrety key exists in list of secrets
func SctExists(s []models.SecretMetadata, sctKy string) bool {
	for _, sct := range s {
		if sct.Key == sctKy {
			return true
		}
	}
	return false
}

// Returns Secret ACL permissions for a given principal
func GetSctACLPermission(s []models.ACLItem, prin string) string {

	for _, acl := range s {
		if acl.Principal == prin {
			return string(acl.Permission)

		}
	}
	return ""
}
