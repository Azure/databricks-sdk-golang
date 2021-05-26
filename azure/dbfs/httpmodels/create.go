package httpmodels

type CreateReq struct {
	Path      string `json:"path,omitempty" url:"path,omitempty"`
	Overwrite bool   `json:"overwrite,omitempty" url:"overwrite,omitempty"`
}

type CreateResp struct {
	Handle int64 `json:"handle,omitempty" url:"handle,omitempty"`
}
