package httpmodels

type PinReq struct {
	ClusterID string `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
}
