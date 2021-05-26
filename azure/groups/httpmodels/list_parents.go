package httpmodels

type ListParentsReq struct {
	UserName  string `json:"user_name,omitempty" url:"user_name,omitempty"`
	GroupName string `json:"group_name,omitempty" url:"group_name,omitempty"`
}

type ListParentsResp struct {
	GroupNames []string `json:"group_names,omitempty" url:"group_names,omitempty"`
}
