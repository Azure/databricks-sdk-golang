package httpmodels

type DeleteReq struct {
	TokenID string `json:"token_id,omitempty" url:"token_id,omitempty"`
}
