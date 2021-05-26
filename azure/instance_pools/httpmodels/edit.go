package httpmodels

//EditReq is to compose the request for modifying an instance pool
type EditReq struct {
	InstancePoolID                     string `json:"instance_pool_id,omitempty" url:"instance_pool_id,omitempty"`
	InstancePoolName                   string `json:"instance_pool_name,omitempty" url:"instance_pool_name,omitempty"`
	MinIdleInstances                   int32  `json:"min_idle_instances,omitempty" url:"min_idle_instances,omitempty"`
	MaxCapacity                        int32  `json:"max_capacity,omitempty" url:"max_capacity,omitempty"`
	NodeTypeID                         string `json:"node_type_id,omitempty" url:"node_type_id,omitempty"`
	IdleInstanceAutoterminationMinutes int32  `json:"idle_instance_autotermination_minutes,omitempty" url:"idle_instance_autotermination_minutes,omitempty"`
}
