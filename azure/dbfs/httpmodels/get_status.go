// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package httpmodels

type GetStatusReq struct {
	Path string `json:"path,omitempty" url:"path,omitempty"`
}

type GetStatusResp struct {
	Path             string `json:"path,omitempty" url:"path,omitempty"`
	IsDir            bool   `json:"is_dir,omitempty" url:"is_dir,omitempty"`
	FileSize         int64  `json:"file_size,omitempty" url:"file_size,omitempty"`
	ModificationTime int64  `json:"modification_time,omitempty" url:"modification_time,omitempty"`
}
