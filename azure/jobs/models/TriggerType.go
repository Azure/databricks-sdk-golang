package models

type TriggerType string

const (
	TriggerTypePeriodic = "PERIODIC"
	TriggerTypeOneTime  = "ONE_TIME"
	TriggerTypeRetry    = "RETRY"
)
