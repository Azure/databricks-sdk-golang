package httpmodels

import (
	"github.com/polar-rams/databricks-sdk-golang/azure/token/models"
)

type CreateResp struct {
	TokenValue string                 `json:"token_value,omitempty" url:"token_value,omitempty"`
	TokenInfo  models.PublicTokenInfo `json:"token_info,omitempty" url:"token_info,omitempty"`
}
