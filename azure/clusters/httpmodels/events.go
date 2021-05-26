// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package httpmodels

import (
	"github.com/Azure/databricks-sdk-golang/azure/clusters/models"
)

type EventsReq struct {
	ClusterID  string                    `json:"cluster_id,omitempty" url:"cluster_id,omitempty"`
	StartTime  int64                     `json:"start_time,omitempty" url:"start_time,omitempty"`
	EndTime    int64                     `json:"end_time,omitempty" url:"end_time,omitempty"`
	Order      models.ListOrder          `json:"order,omitempty" url:"order,omitempty"`
	EventTypes []models.ClusterEventType `json:"event_types,omitempty" url:"event_types,omitempty"`
	Offset     int64                     `json:"offset,omitempty" url:"offset,omitempty"`
	Limit      int64                     `json:"limit,omitempty" url:"limit,omitempty"`
}

type EventsResp struct {
	Events     *[]models.ClusterEvent `json:"events,omitempty" url:"events,omitempty"`
	NextPage   EventsReq              `json:"next_page,omitempty" url:"next_page,omitempty"`
	TotalCount int32                  `json:"total_count,omitempty" url:"total_count,omitempty"`
}
