// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package azure_test

import (
	"testing"

	"github.com/Azure/databricks-sdk-golang/azure/ip_access_lists/httpmodels"
	"github.com/Azure/databricks-sdk-golang/azure/ip_access_lists/models"
	"github.com/stretchr/testify/assert"
)

func TestAzureIPAccessListsSetCheckEnabled(t *testing.T) {
	// Disable
	e := c.IPAccessLists().Set(false)
	assert.Nil(t, e, "failed to disable IP access lists")

	// Check
	respCheckEnabled, _ := c.IPAccessLists().CheckEnabled()
	assert.Nil(t, e, "failed to check IP access lists enabled")
	assert.Equal(t, "false", respCheckEnabled.EnableIpAccessLists, "IP access lists should be disabled")

	// Reset
	_ = c.IPAccessLists().Set(true)
}

func TestAzureIPAccessListsAddGetAllGetDelete(t *testing.T) {
	// Add
	reqAdd := httpmodels.AddReq{
		Label:    "Test IP Access List",
		ListType: models.ListTypeBlock,
		IPAddresses: &[]string{
			"8.8.8.8",
		},
	}
	respAdd, e := c.IPAccessLists().Add(reqAdd)
	assert.Nil(t, e, "failed to add a new IP access list")
	assert.NotNil(t, respAdd.IPAccessList)
	assert.Equal(t, reqAdd.Label, respAdd.IPAccessList.Label)
	assert.Equal(t, reqAdd.ListType, respAdd.IPAccessList.ListType)
	assert.Equal(t, reqAdd.IPAddresses, respAdd.IPAccessList.IPAddresses)

	// GetAll
	respGetAll, e := c.IPAccessLists().GetAll()
	assert.Nil(t, e, "failed to get all IP access lists")
	assert.NotNil(t, respGetAll.IPAccessLists)
	assert.Equal(t, 1, len(respGetAll.IPAccessLists))

	// Get
	respGet, e := c.IPAccessLists().Get(respAdd.IPAccessList.ListID)
	assert.Nil(t, e, "failed to get IP access list")
	assert.NotNil(t, respGet.IPAccessList)
	assert.Equal(t, respAdd.IPAccessList, respGet.IPAccessList)

	// Reset
	e = c.IPAccessLists().Delete(respAdd.IPAccessList.ListID)
	assert.Nil(t, e, "failed to delete IP access list")
}

func TestAzureIPAccessListsReplaceGetDelete(t *testing.T) {
	// Add
	reqAdd := httpmodels.AddReq{
		Label:    "Test IP Access List Add",
		ListType: models.ListTypeBlock,
		IPAddresses: &[]string{
			"8.8.8.8",
		},
	}
	respAdd, e := c.IPAccessLists().Add(reqAdd)
	assert.Nil(t, e, "failed to add a new IP access list")
	assert.Equal(t, reqAdd.Label, respAdd.IPAccessList.Label)
	assert.Equal(t, reqAdd.ListType, respAdd.IPAccessList.ListType)
	assert.Equal(t, reqAdd.IPAddresses, respAdd.IPAccessList.IPAddresses)

	// Replace
	reqReplace := httpmodels.ReplaceReq{
		ListID:   respAdd.IPAccessList.ListID,
		Label:    "Test IP Access List Replace",
		ListType: models.ListTypeBlock,
		IPAddresses: &[]string{
			"1.1.1.1",
		},
		Enabled: true,
	}
	respReplace, e := c.IPAccessLists().Replace(respAdd.IPAccessList.ListID, reqReplace)
	assert.Nil(t, e, "failed to replace IP access list")
	assert.Equal(t, reqReplace.ListID, respReplace.IPAccessList.ListID)
	assert.Equal(t, reqReplace.Label, respReplace.IPAccessList.Label)
	assert.Equal(t, reqReplace.ListType, respReplace.IPAccessList.ListType)
	assert.Equal(t, reqReplace.IPAddresses, respReplace.IPAccessList.IPAddresses)
	assert.Equal(t, reqReplace.Enabled, respReplace.IPAccessList.Enabled)

	// Get
	respGet, e := c.IPAccessLists().Get(respAdd.IPAccessList.ListID)
	assert.Nil(t, e, "failed to get IP access list")
	assert.NotNil(t, respGet)
	assert.Equal(t, respReplace.IPAccessList, respGet.IPAccessList)

	// Reset
	_ = c.IPAccessLists().Delete(respAdd.IPAccessList.ListID)
	assert.Nil(t, e, "failed to delete IP access list")
}

func TestAzureIPAccessListsUpdateGetDelete(t *testing.T) {
	// Add
	reqAdd := httpmodels.AddReq{
		Label:    "Test IP Access List",
		ListType: models.ListTypeBlock,
		IPAddresses: &[]string{
			"8.8.8.8",
		},
	}
	respAdd, e := c.IPAccessLists().Add(reqAdd)
	assert.Nil(t, e, "failed to add a new IP access list")
	assert.Equal(t, reqAdd.Label, respAdd.IPAccessList.Label)
	assert.Equal(t, reqAdd.ListType, respAdd.IPAccessList.ListType)
	assert.Equal(t, reqAdd.IPAddresses, respAdd.IPAccessList.IPAddresses)

	// Update
	reqUpdate := httpmodels.UpdateReq{
		Label: "Test IP Access List Update",
	}
	respUpdate, e := c.IPAccessLists().Update(respAdd.IPAccessList.ListID, reqUpdate)
	assert.Nil(t, e, "failed to update IP access list")
	assert.Equal(t, reqUpdate.Label, respUpdate.IPAccessList.Label)

	reqUpdate = httpmodels.UpdateReq{
		IPAddresses: &[]string{
			"8.8.8.8",
			"2.2.2.2",
		},
	}
	respUpdate, e = c.IPAccessLists().Update(respAdd.IPAccessList.ListID, reqUpdate)
	assert.Nil(t, e, "failed to update IP access list")
	assert.Equal(t, reqUpdate.IPAddresses, respUpdate.IPAccessList.IPAddresses)

	// Get
	respGet, e := c.IPAccessLists().Get(respAdd.IPAccessList.ListID)
	assert.Nil(t, e, "failed to get IP access list")
	assert.NotNil(t, respGet)
	assert.Equal(t, respUpdate.IPAccessList, respGet.IPAccessList)

	// Reset
	_ = c.IPAccessLists().Delete(respAdd.IPAccessList.ListID)
	assert.Nil(t, e, "failed to delete IP access list")
}
