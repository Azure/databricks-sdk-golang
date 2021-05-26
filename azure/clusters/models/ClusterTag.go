// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package models

type ClusterTag struct {
	Key   string `json:"key,omitempty" url:"key,omitempty"`
	Value string `json:"value,omitempty" url:"value,omitempty"`
}
