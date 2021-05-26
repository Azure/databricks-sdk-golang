package models

type InstancePoolStats struct {
	UsedCount        int32 `json:"used_count,omitempty" url:"used_count,omitempty"`
	IdleCount        int32 `json:"idle_count,omitempty" url:"idle_count,omitempty"`
	PendingUsedCount int32 `json:"pending_used_count,omitempty" url:"pending_used_count,omitempty"`
	PendingIdleCount int32 `json:"pending_idle_count,omitempty" url:"pending_idle_count,omitempty"`
}
