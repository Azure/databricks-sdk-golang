// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package httpmodels

import "github.com/Azure/databricks-sdk-golang/azure/jobs/models"

type RunsGetReq struct {
	RunID int64 `json:"run_id,omitempty" url:"run_id,omitempty"`
}

type RunsGetResp struct {
	JobID                int64                  `json:"job_id,omitempty" url:"job_id,omitempty"`
	RunID                int64                  `json:"run_id,omitempty" url:"run_id,omitempty"`
	NumberInJob          int64                  `json:"number_in_job,omitempty" url:"number_in_job,omitempty"`
	OriginalAttemptRunID int64                  `json:"original_attempt_run_id,omitempty" url:"original_attempt_run_id,omitempty"`
	State                models.RunState        `json:"state,omitempty" url:"state,omitempty"`
	Schedule             models.CronSchedule    `json:"schedule,omitempty" url:"schedule,omitempty"`
	Task                 *models.JobTask        `json:"task,omitempty" url:"task,omitempty"`
	ClusterSpec          *models.ClusterSpec    `json:"cluster_spec,omitempty" url:"cluster_spec,omitempty"`
	ClusterInstance      models.ClusterInstance `json:"cluster_instance,omitempty" url:"cluster_instance,omitempty"`
	OverridingParameters *models.RunParameters  `json:"overriding_parameters,omitempty" url:"overriding_parameters,omitempty"`
	StartTime            int64                  `json:"start_time,omitempty" url:"start_time,omitempty"`
	EndTime              int64                  `json:"end_time,omitempty" url:"end_time,omitempty"`
	SetupDuration        int64                  `json:"setup_duration,omitempty" url:"setup_duration,omitempty"`
	ExecutionDuration    int64                  `json:"execution_duration,omitempty" url:"execution_duration,omitempty"`
	CleanupDuration      int64                  `json:"cleanup_duration,omitempty" url:"cleanup_duration,omitempty"`
	Trigger              models.TriggerType     `json:"trigger,omitempty" url:"trigger,omitempty"`
	CreatorUserName      string                 `json:"creator_user_name,omitempty" url:"creator_user_name,omitempty"`
	RunPageURL           string                 `json:"run_page_url,omitempty" url:"run_page_url,omitempty"`
}
