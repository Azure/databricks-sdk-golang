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
	ListID       string          `json:"list_id,omitempty" url:"list_id,omitempty"`
	Label        string          `json:"label,omitempty" url:"label,omitempty"`
	IPAddresses  *[]string       `json:"ip_addresses,omitempty" url:"ip_addresses,omitempty"`
	AddressCount int             `json:"address_count,omitempty" url:"address_count,omitempty"`
	ListType     models.ListType `json:"list_type,omitempty" url:"list_type,omitempty"`
	CreatedAt    int             `json:"created_at,omitempty" url:"created_at,omitempty"`
	CreatedBy    int             `json:"created_by,omitempty" url:"created_by,omitempty"`
	UpdatedAt    int             `json:"updated_at,omitempty" url:"updated_at,omitempty"`
	UpdatedBy    int             `json:"updated_by,omitempty" url:"updated_by,omitempty"`
	Enabled      bool            `json:"enabled,omitempty" url:"enabled,omitempty"`
}
