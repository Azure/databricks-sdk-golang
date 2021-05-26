// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package models

type DiskSpec struct {
	DiskType  DiskType `json:"disk_type,omitempty" url:"disk_type,omitempty"`
	DiskCount int32    `json:"disk_count,omitempty" url:"disk_count,omitempty"`
	DiskSize  int32    `json:"disk_size,omitempty" url:"disk_size,omitempty"`
}
