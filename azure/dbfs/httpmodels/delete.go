package httpmodels

type DeleteReq struct {
	Path      string `json:"path,omitempty" url:"path,omitempty"`
	Recursive bool   `json:"recursive,omitempty" url:"recursive,omitempty"`
}