package httpmodels

import (
	"github.com/polar-rams/databricks-sdk-golang/azure/token/models"
)

type ListResp struct {
	TokenInfos *[]models.PublicTokenInfo `json:"token_infos,omitempty" url:"token_infos,omitempty"`
}
