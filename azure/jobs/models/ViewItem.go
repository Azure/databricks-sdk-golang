// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package models

type ViewItem struct {
	Content string   `json:"content,omitempty" url:"content,omitempty"`
	Name    string   `json:"name,omitempty" url:"name,omitempty"`
	Type    ViewType `json:"type,omitempty" url:"type,omitempty"`
}
