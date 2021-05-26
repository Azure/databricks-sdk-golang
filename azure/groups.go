package azure

import (
	"encoding/json"
	"net/http"

	"github.com/polar-rams/databricks-sdk-golang/azure/groups/httpmodels"
)

// GroupsAPI exposes the Groups API
type GroupsAPI struct {
	Client DBClient
}

func (a GroupsAPI) init(client DBClient) GroupsAPI {
	a.Client = client
	return a
}

// AddMember adds a user or group to a group
func (a GroupsAPI) AddMember(req httpmodels.AddMemberReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/groups/add-member", req, nil)
	return err
}

// Create creates a new group with the given name
func (a GroupsAPI) Create(req httpmodels.CreateReq) (httpmodels.CreateResp, error) {
	var resp httpmodels.CreateResp

	jsonResp, err := a.Client.performQuery(http.MethodPost, "/groups/create", req, nil)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(jsonResp, &resp)
	return resp, err
}

// ListMembers returns all of the members of a particular group
func (a GroupsAPI) ListMembers(req httpmodels.ListMembersReq) (httpmodels.ListMembersResp, error) {
	var resp httpmodels.ListMembersResp

	jsonResp, err := a.Client.performQuery(http.MethodGet, "/groups/list-members", req, nil)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(jsonResp, &resp)
	return resp, err
}

// List returns all of the groups in an organization
func (a GroupsAPI) List() (httpmodels.ListResp, error) {
	var resp httpmodels.ListResp

	jsonResp, err := a.Client.performQuery(http.MethodGet, "/groups/list", nil, nil)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(jsonResp, &resp)
	return resp, err
}

// ListParents retrieves all groups in which a given user or group is a member
func (a GroupsAPI) ListParents(listParentsReq httpmodels.ListParentsReq) (httpmodels.ListParentsResp, error) {
	var resp httpmodels.ListParentsResp

	jsonResp, err := a.Client.performQuery(http.MethodGet, "/groups/list-parents", listParentsReq, nil)
	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(jsonResp, &resp)
	return resp, err
}

// RemoveMember removes a user or group from a group
func (a GroupsAPI) RemoveMember(removeMemberReq httpmodels.RemoveMemberReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/groups/remove-member", removeMemberReq, nil)
	return err
}

// Delete removes a group from this organization
func (a GroupsAPI) Delete(deleteReq httpmodels.DeleteReq) error {
	_, err := a.Client.performQuery(http.MethodPost, "/groups/delete", deleteReq, nil)
	return err
}
