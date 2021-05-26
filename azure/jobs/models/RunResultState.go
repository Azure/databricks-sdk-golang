// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package models

type RunResultState string

const (
	RunResultStateSuccess  = "SUCCESS"
	RunResultStateFailed   = "FAILED"
	RunResultStateTimedout = "TIMEDOUT"
	RunResultStateCanceled = "CANCELED"
)
