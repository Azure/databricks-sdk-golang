package httpmodels

type StartReq struct {
	ClusterID string `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
}
