package httpmodels

type CloseReq struct {
	Handle int64 `json:"handle,omitempty" url:"handle,omitempty"`
}