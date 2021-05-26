// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package models

type LibraryFullStatus struct {
	Library                 *Library             `json:"library,omitempty" url:"library,omitempty"`
	Status                  LibraryInstallStatus `json:"status,omitempty" url:"status,omitempty"`
	Messages                *[]string            `json:"messages,omitempty" url:"messages,omitempty"`
	IsLibraryForAllClusters bool                 `json:"is_library_for_all_clusters,omitempty" url:"is_library_for_all_clusters,omitempty"`
}
