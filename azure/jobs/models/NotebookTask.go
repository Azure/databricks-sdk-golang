// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package models

type NotebookTask struct {
	NotebookPath      string             `json:"notebook_path,omitempty" url:"notebook_path,omitempty"`
	RevisionTimestamp int64              `json:"revision_timestamp,omitempty" url:"revision_timestamp,omitempty"`
	BaseParameters    *map[string]string `json:"base_parameters,omitempty" url:"base_parameters,omitempty"`
}
