// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package models

type ClusterEventType string

const (
	ClusterEventTypeCreating               = "CREATING"
	ClusterEventTypeDidNotExpandDisk       = "DID_NOT_EXPAND_DISK"
	ClusterEventTypeExpandedDisk           = "EXPANDED_DISK"
	ClusterEventTypeFailedToExpandDisk     = "FAILED_TO_EXPAND_DISK"
	ClusterEventTypeInitScriptStarting     = "INIT_SCRIPTS_STARTING"
	ClusterEventTypeInitScriptFinished     = "INIT_SCRIPTS_FINISHED"
	ClusterEventTypeStarting               = "STARTING"
	ClusterEventTypeRestarting             = "RESTARTING"
	ClusterEventTypeTerminating            = "TERMINATING"
	ClusterEventTypeEdited                 = "EDITED"
	ClusterEventTypeRunning                = "RUNNING"
	ClusterEventTypeResizing               = "RESIZING"
	ClusterEventTypeUpsizeCompleted        = "UPSIZE_COMPLETED"
	ClusterEventTypeNodesLost              = "NODES_LOST"
	ClusterEventTypeDriverHealthy          = "DRIVER_HEALTHY"
	ClusterEventTypeDriverUnavailable      = "DRIVER_UNAVAILABLE"
	ClusterEventTypeSparkException         = "SPARK_EXCEPTION"
	ClusterEventTypeDriverNotResponding    = "DRIVER_NOT_RESPONDING"
	ClusterEventTypeDbfsDown               = "DBFS_DOWN"
	ClusterEventTypeMetastoreDown          = "METASTORE_DOWN"
	ClusterEventTypeAutoscalingStatsReport = "AUTOSCALING_STATS_REPORT"
)
