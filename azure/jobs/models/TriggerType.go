// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package models

type TriggerType string

const (
	TriggerTypePeriodic = "PERIODIC"
	TriggerTypeOneTime  = "ONE_TIME"
	TriggerTypeRetry    = "RETRY"
)
