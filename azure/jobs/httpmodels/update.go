package httpmodels

import "github.com/polar-rams/databricks-sdk-golang/azure/jobs/models"

type UpdateReq struct {
	JobID          int64               `json:"job_id,omitempty" url:"job_id,omitempty"`
	NewSettings    *models.JobSettings `json:"new_settings,omitempty" url:"new_settings,omitempty"`
	FieldsToRemove *[]string           `json:"fields_to_remove,omitempty" url:"fields_to_remove,omitempty"`
}
