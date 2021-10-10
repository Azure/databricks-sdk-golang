package httpmodels

import (
	"github.com/Azure/databricks-sdk-golang/azure/ip_access_lists/models"
)

type GetResp struct {
	IPAccessList models.IPAccessList `json:"ip_access_list,omitempty" url:"ip_access_list,omitempty"`
}
