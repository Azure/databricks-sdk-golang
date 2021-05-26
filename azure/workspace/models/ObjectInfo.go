// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package models

type ObjectInfo struct {
	ObjectType ObjectType `json:"object_type,omitempty" url:"object_type,omitempty"`
	ObjectID   int64      `json:"object_id,omitempty" url:"object_id,omitempty"`
	Path       string     `json:"path,omitempty" url:"path,omitempty"`
	Language   Language   `json:"language,omitempty" url:"language,omitempty"`
}
