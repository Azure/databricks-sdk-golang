package models

type ResizeCause string

const (
	ResizeCauseAutoscale    = "AUTOSCALE"
	ResizeCauseUserRequest  = "USER_REQUEST"
	ResizeCauseAutorecovery = "AUTORECOVERY"
)
