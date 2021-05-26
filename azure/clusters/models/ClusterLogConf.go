// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package models

type ClusterLogConf struct {
	Dbfs DbfsStorageInfo `json:"dbfs,omitempty" url:"dbfs,omitempty"`
}
