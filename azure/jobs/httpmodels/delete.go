package httpmodels

type DeleteReq struct {
	JobID int64 `json:"job_id,omitempty" url:"job_id,omitempty"`
}
