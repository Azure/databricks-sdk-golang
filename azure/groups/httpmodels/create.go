package httpmodels

type CreateReq struct {
	GroupName string `json:"group_name,omitempty" url:"group_name,omitempty"`
}

type CreateResp struct {
	GroupName string `json:"group_name,omitempty" url:"group_name,omitempty"`
}
