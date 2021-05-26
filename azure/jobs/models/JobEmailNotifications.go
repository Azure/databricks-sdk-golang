package models

type JobEmailNotifications struct {
	OnStart               *[]string `json:"on_start,omitempty" url:"on_start,omitempty"`
	OnSuccess             *[]string `json:"on_success,omitempty" url:"on_success,omitempty"`
	OnFailure             *[]string `json:"on_failure,omitempty" url:"on_failure,omitempty"`
	NoAlertForSkippedRuns bool      `json:"no_alert_for_skipped_runs,omitempty" url:"no_alert_for_skipped_runs,omitempty"`
}
