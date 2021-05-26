package httpmodels

type MkdirsReq struct {
	Path string `json:"path,omitempty" url:"path,omitempty"`
}
