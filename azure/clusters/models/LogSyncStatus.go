// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package models

type LogSyncStatus struct {
	LastAttempted int64  `json:"last_attempted,omitempty" url:"last_attempted,omitempty"`
	LastException string `json:"last_exception,omitempty" url:"last_exception,omitempty"`
}
