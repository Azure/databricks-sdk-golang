// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package models

type CronSchedule struct {
	QuartzCronExpression string `json:"quartz_cron_expression,omitempty" url:"quartz_cron_expression,omitempty"`
	TimezoneID           string `json:"timezone_id,omitempty" url:"timezone_id,omitempty"`
	PauseStatus          string `json:"pause_status,omitempty" url:"pause_status,omitempty"`
}
