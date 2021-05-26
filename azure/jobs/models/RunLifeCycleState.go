// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package models

type RunLifeCycleState string

const (
	RunLifeCycleStatePending       = "PENDING"
	RunLifeCycleStateRunning       = "RUNNING"
	RunLifeCycleStateTerminating   = "TERMINATING"
	RunLifeCycleStateTerminated    = "TERMINATED"
	RunLifeCycleStateSkipped       = "SKIPPED"
	RunLifeCycleStateInternalError = "INTERNAL_ERROR"
)
