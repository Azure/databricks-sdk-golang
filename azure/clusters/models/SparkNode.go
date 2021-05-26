// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package models

type SparkNode struct {
	PrivateIP      string `json:"private_ip,omitempty" url:"private_ip,omitempty"`
	PublicDNS      string `json:"public_dns,omitempty" url:"public_dns,omitempty"`
	NodeID         string `json:"node_id,omitempty" url:"node_id,omitempty"`
	InstanceID     string `json:"instance_id,omitempty" url:"instance_id,omitempty"`
	StartTimestamp int64  `json:"start_timestamp,omitempty" url:"start_timestamp,omitempty"`
	HostPrivateIP  string `json:"host_private_ip,omitempty" url:"host_private_ip,omitempty"`
}
