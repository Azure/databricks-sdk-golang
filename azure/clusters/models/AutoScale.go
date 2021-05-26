// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package models

type AutoScale struct {
	MinWorkers int32 `json:"min_workers,omitempty" url:"min_workers,omitempty"`
	MaxWorkers int32 `json:"max_workers,omitempty" url:"max_workers,omitempty"`
}
