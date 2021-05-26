package httpmodels

type ListResp struct {
	GroupNames []string `json:"group_names,omitempty" url:"group_names,omitempty"`
}
