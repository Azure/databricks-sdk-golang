package httpmodels

import (
	"github.com/Azure/databricks-sdk-golang/azure/ip_access_lists/models"
)

type GetAllResp struct {
	IPAccessLists []models.IPAccessList `json:"ip_access_lists,omitempty" url:"ip_access_lists,omitempty"`
}
