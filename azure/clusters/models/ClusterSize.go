package models

type ClusterSize struct {
	NumWorkers int32     `json:"num_workers,omitempty" url:"num_workers,omitempty"`
	Autoscale  AutoScale `json:"autoscale,omitempty" url:"autoscale,omitempty"`
}
