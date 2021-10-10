package httpmodels

import (
	"github.com/Azure/databricks-sdk-golang/azure/ip_access_lists/models"
)

type SetReq struct {
	EnableIpAccessLists models.EnableIpAccessLists `json:"enableIpAccessLists,omitempty" url:"enableIpAccessLists,omitempty"`
}
