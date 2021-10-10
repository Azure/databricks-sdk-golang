package httpmodels

import "github.com/Azure/databricks-sdk-golang/azure/ip_access_lists/models"

type ReplaceReq struct {
	ListID      string          `json:"list_id,omitempty" url:"list_id,omitempty"`
	Label       string          `json:"label,omitempty" url:"label,omitempty"`
	ListType    models.ListType `json:"list_type,omitempty" url:"list_type,omitempty"`
	IPAddresses *[]string       `json:"ip_addresses,omitempty" url:"ip_addresses,omitempty"`
	Enabled     bool            `json:"enabled,omitempty" url:"enabled,omitempty"`
}

type ReplaceResp struct {
	IPAccessList models.IPAccessList `json:"ip_access_list,omitempty" url:"ip_access_list,omitempty"`
}
