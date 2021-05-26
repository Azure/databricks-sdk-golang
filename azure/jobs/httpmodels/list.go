package httpmodels

import (
	"github.com/polar-rams/databricks-sdk-golang/azure/jobs/models"
)

type ListResp struct {
	Jobs *[]models.Job `json:"jobs,omitempty" url:"jobs,omitempty"`
}
