package models

type TerminationType string

const (
	Success      = "SUCCESS"
	ClientError  = "CLIENT_ERROR"
	ServiceFault = "SERVICE_FAULT"
	CloudFailure = "CLOUD_FAILURE"
)
