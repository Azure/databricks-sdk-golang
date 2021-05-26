// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package models

type ClusterState string

const (
	ClusterStatePending     = "PENDING"
	ClusterStateRunning     = "RUNNING"
	ClusterStateRestarting  = "RESTARTING"
	ClusterStateResizing    = "RESIZING"
	ClusterStateTerminating = "TERMINATING"
	ClusterStateError       = "ERROR"
	ClusterStateUnknown     = "UNKNOWN"
	ClusterStateTERMINATED  = "TERMINATED"
)
