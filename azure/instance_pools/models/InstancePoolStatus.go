// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package models

type InstancePoolStatus struct {
	PendingInstanceErrors *[]PendingInstanceError `json:"pending_instance_errors,omitempty" url:"pending_instance_errors,omitempty"`
}
