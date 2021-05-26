package azure_test

import (
	"fmt"
	"testing"

	"github.com/polar-rams/databricks-sdk-golang/azure/groups/httpmodels"
	"github.com/polar-rams/databricks-sdk-golang/azure/groups/models"
	"github.com/stretchr/testify/assert"
)

const (
	pfxGroup    = "testGroup"
	adminsGroup = "admins"
)

func TestAzureGroupsCreateListDelete(t *testing.T) {

	groupName := fmt.Sprintf("%s%s", pfxGroup, randSeq(6))

	// Test create group and assert no errors
	createReq := httpmodels.CreateReq{
		GroupName: groupName,
	}
	createRes, e := c.Groups().Create(createReq)
	assert.Nil(t, e, fmt.Sprintf("failed to create group %+v", createReq))
	assert.NotEmpty(t, createRes.GroupName, fmt.Sprintf("failed to create group %+v", createReq))

	// Test list groups and assert group exists
	listRes, e := c.Groups().List()
	assert.Nil(t, e, "failed to list groups")
	assert.True(t, GroupExists(listRes.GroupNames, groupName), fmt.Sprintf("created group not found %+v", createReq))

	// Test delete group and assert no errors
	deleteReq := httpmodels.DeleteReq{
		GroupName: groupName,
	}
	e = c.Groups().Delete(deleteReq)
	assert.Nil(t, e, fmt.Sprintf("failed to delete group %+v", deleteReq))

	// Test list groups and assert group does not exist
	listRes, e = c.Groups().List()
	assert.Nil(t, e, "failed to list groups")
	assert.False(t, GroupExists(listRes.GroupNames, groupName), fmt.Sprintf("deleted group still exists %+v", deleteReq))
}

func TestAzureGroupMembersAddListRemoveUser(t *testing.T) {

	// Setup: List members in a default 'users' group, assess no errors and pick one user for further tests purposes
	listAdminReq := httpmodels.ListMembersReq{
		GroupName: adminsGroup,
	}
	listAdminResp, e := c.Groups().ListMembers(listAdminReq)
	assert.Nil(t, e, fmt.Sprintf("failed to list members in a group %s", adminsGroup))

	hasAdmins := listAdminResp.Members != nil && len(*listAdminResp.Members) > 0
	assert.True(t, e != nil || hasAdmins, fmt.Sprintf("no members in a group %s", adminsGroup))
	if !hasAdmins {
		return
	}

	userName := (*listAdminResp.Members)[0].UserName

	// Setup: create parent group
	parentGroup := fmt.Sprintf("%s%s", pfxGroup, randSeq(6))

	createGroupReq := httpmodels.CreateReq{
		GroupName: parentGroup,
	}
	_, e = c.Groups().Create(createGroupReq)
	assert.Nil(t, e, fmt.Sprintf("failed to create group %+v", createGroupReq))

	// Add member to a group and assert no errors
	addMemberReq := httpmodels.AddMemberReq{
		UserName:   userName,
		ParentName: parentGroup,
	}
	e = c.Groups().AddMember(addMemberReq)
	assert.Nil(t, e, fmt.Sprintf("failed to add member to a group %+v", addMemberReq))

	// List members in a group and assert member added
	listMemberReq := httpmodels.ListMembersReq{
		GroupName: parentGroup,
	}
	listMemberResp, e := c.Groups().ListMembers(listMemberReq)
	assert.Nil(t, e, fmt.Sprintf("failed to list members in a group %s", parentGroup))
	assert.True(t, GroupMemberExists(listMemberResp.Members, userName, ""), fmt.Sprintf("added member %s not found in a group %s", userName, parentGroup))

	// List member's parents and assert parent group present
	listParentReq := httpmodels.ListParentsReq{
		UserName: userName,
	}
	listParentResp, e := c.Groups().ListParents(listParentReq)
	assert.Nil(t, e, fmt.Sprintf("failed to list parents for a member %s", parentGroup))
	assert.True(t, GroupExists(listParentResp.GroupNames, parentGroup), fmt.Sprintf("parent group %s not listed for a member %s", parentGroup, userName))

	// Remove member from a group and assert no errors
	removeMemberReq := httpmodels.RemoveMemberReq{
		UserName:   userName,
		ParentName: parentGroup,
	}
	e = c.Groups().RemoveMember(removeMemberReq)
	assert.Nil(t, e, fmt.Sprintf("failed to remove member from a group %+v", removeMemberReq))

	// List members in a group and assert member removed
	listMemberResp, e = c.Groups().ListMembers(listMemberReq)
	assert.Nil(t, e, fmt.Sprintf("failed to list members in a group %s", parentGroup))
	assert.False(t, GroupMemberExists(listMemberResp.Members, userName, ""), fmt.Sprintf("removed member still exists in a group %+v", removeMemberReq))

	// Cleanup: delete parent group
	deleteReq := httpmodels.DeleteReq{
		GroupName: parentGroup,
	}
	e = c.Groups().Delete(deleteReq)
	assert.Nil(t, e, fmt.Sprintf("failed to delete group %+v", deleteReq))
}

func TestAzureGroupMembersAddListRemoveGroup(t *testing.T) {
	// Setup: create parent group
	parentGroup := fmt.Sprintf("%s%s", pfxGroup, randSeq(6))

	createGroupReq := httpmodels.CreateReq{
		GroupName: parentGroup,
	}
	_, e := c.Groups().Create(createGroupReq)
	assert.Nil(t, e, fmt.Sprintf("failed to create parent group %+v", createGroupReq))

	// Setup: create member group
	memberGroup := fmt.Sprintf("%s%s", pfxGroup, randSeq(6))

	createGroupReq = httpmodels.CreateReq{
		GroupName: memberGroup,
	}
	_, e = c.Groups().Create(createGroupReq)
	assert.Nil(t, e, fmt.Sprintf("failed to create member group %+v", createGroupReq))

	// Add member to a group and assert no errors
	addMemberReq := httpmodels.AddMemberReq{
		GroupName:  memberGroup,
		ParentName: parentGroup,
	}
	e = c.Groups().AddMember(addMemberReq)
	assert.Nil(t, e, fmt.Sprintf("failed to add member to a group %+v", addMemberReq))

	// List members in a group and assert member added
	listMemberReq := httpmodels.ListMembersReq{
		GroupName: parentGroup,
	}
	listMemberResp, e := c.Groups().ListMembers(listMemberReq)
	assert.Nil(t, e, fmt.Sprintf("failed to list members in a group %s", parentGroup))
	assert.True(t, GroupMemberExists(listMemberResp.Members, "", memberGroup), fmt.Sprintf("added member %s not found in a group %s", memberGroup, parentGroup))

	// List member's parents and assert parent group present
	listParentReq := httpmodels.ListParentsReq{
		GroupName: memberGroup,
	}
	listParentResp, e := c.Groups().ListParents(listParentReq)
	assert.Nil(t, e, fmt.Sprintf("failed to list parents for a member %s", parentGroup))
	assert.True(t, GroupExists(listParentResp.GroupNames, parentGroup), fmt.Sprintf("parent group %s not listed for a member %s", parentGroup, memberGroup))

	// Remove member from a group and assert no errors
	removeMemberReq := httpmodels.RemoveMemberReq{
		GroupName:  memberGroup,
		ParentName: parentGroup,
	}
	e = c.Groups().RemoveMember(removeMemberReq)
	assert.Nil(t, e, fmt.Sprintf("failed to remove member from a group %+v", removeMemberReq))

	// List members in a group and assert member removed
	listMemberResp, e = c.Groups().ListMembers(listMemberReq)
	assert.Nil(t, e, fmt.Sprintf("failed to list members in a group %s", parentGroup))
	assert.False(t, GroupMemberExists(listMemberResp.Members, "", memberGroup), fmt.Sprintf("removed member still exists in a group %+v", removeMemberReq))

	// Cleanup: delete parent group
	deleteReq := httpmodels.DeleteReq{
		GroupName: parentGroup,
	}
	e = c.Groups().Delete(deleteReq)
	assert.Nil(t, e, fmt.Sprintf("failed to delete parent group %+v", deleteReq))

	// Cleanup: delete member group
	deleteReq = httpmodels.DeleteReq{
		GroupName: memberGroup,
	}
	e = c.Groups().Delete(deleteReq)
	assert.Nil(t, e, fmt.Sprintf("failed to delete member group %+v", deleteReq))
}

// GroupExists checks if group exists in the list of group names
func GroupExists(groups []string, groupName string) bool {
	for _, g := range groups {
		if g == groupName {
			return true
		}
	}
	return false
}

// GroupMemberExists checks if user exists in the list of group members
func GroupMemberExists(members *[]models.PrincipalName, userName string, groupName string) bool {
	if members != nil {
		for _, m := range *members {
			if len(userName) > 0 && m.UserName == userName {
				return true
			}
			if len(groupName) > 0 && m.GroupName == groupName {
				return true
			}
		}
	}
	return false
}
