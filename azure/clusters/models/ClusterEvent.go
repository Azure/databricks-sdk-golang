// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package models

type ClusterEvent struct {
	ClusterID string           `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
	Timestamp int64            `json:"timestamp,omitempty" url:"timestamp,omitempty"`
	Type      ClusterEventType `json:"type,omitempty" url:"type,omitempty"`
	Details   *EventDetails    `json:"details,omitempty" url:"details,omitempty"`
}
