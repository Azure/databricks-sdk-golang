package models

type TerminationCode string

const (
	TerminationCodeUserRequest                = "USER_REQUEST"
	TerminationCodeJobFinished                = "JOB_FINISHED"
	TerminationCodeInactivity                 = "INACTIVITY"
	TerminationCodeCloudProviderShutdown      = "CLOUD_PROVIDER_SHUTDOWN"
	TerminationCodeCommunicationLost          = "COMMUNICATION_LOST"
	TerminationCodeCloudProviderLaunchFailure = "CLOUD_PROVIDER_LAUNCH_FAILURE"
	TerminationCodeSparkStartupFailure        = "SPARK_STARTUP_FAILURE"
	TerminationCodeInvalidArgument            = "INVALID_ARGUMENT"
	TerminationCodeUnexpectedLaunchFailure    = "UNEXPECTED_LAUNCH_FAILURE"
	TerminationCodeInternalError              = "INTERNAL_ERROR"
	TerminationCodeInstanceUnreachable        = "INSTANCE_UNREACHABLE"
	TerminationCodeRequestRejected            = "REQUEST_REJECTED"
	TerminationCodeInitScriptFailure          = "INIT_SCRIPT_FAILURE"
	TerminationCodeTrialExpired               = "TRIAL_EXPIRED"
)
