// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package azure_test

import (
	"fmt"
	"testing"

	"github.com/Azure/databricks-sdk-golang/azure/token/httpmodels"
	"github.com/Azure/databricks-sdk-golang/azure/token/models"
	"github.com/stretchr/testify/assert"
)

const (
	pfxToken = "testToken"
)

func TestAzureTokenCreateListRevoke(t *testing.T) {

	comment := fmt.Sprintf("%s%s", pfxToken, randSeq(6))

	// Test create token and assert no errors
	createReq := httpmodels.CreateReq{
		Comment:         comment,
		LifetimeSeconds: 300,
	}
	createRes, e := c.Token().Create(createReq)
	assert.Nil(t, e, fmt.Sprintf("failed to create token %+v", createReq))
	assert.NotNil(t, createRes.TokenInfo, fmt.Sprintf("failed to create token %+v", createReq))

	// Test list token and assert token exists
	listRes, e := c.Token().List()
	assert.Nil(t, e, "failed to list tokens")
	assert.True(t, TokenExists(listRes.TokenInfos, comment), fmt.Sprintf("created token not found %+v", createReq))

	// Test delete token and assert no errors
	deleteReq := httpmodels.DeleteReq{
		TokenID: createRes.TokenInfo.TokenID,
	}
	e = c.Token().Revoke(deleteReq)
	assert.Nil(t, e, fmt.Sprintf("failed to delete token %+v", deleteReq))

	// Test list token and assert token does not exist
	listRes, e = c.Token().List()
	assert.Nil(t, e, "failed to list tokens")
	assert.False(t, TokenExists(listRes.TokenInfos, comment), fmt.Sprintf("deleted token still exists %+v", deleteReq))
}

// TokenExists checks if token exists in the list of tokens
func TokenExists(tokens *[]models.PublicTokenInfo, comment string) bool {
	if tokens != nil {
		for _, t := range *tokens {
			if t.Comment == comment {
				return true
			}
		}
	}
	return false
}
