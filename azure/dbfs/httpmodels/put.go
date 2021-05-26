// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package httpmodels

type PutReq struct {
	Path      string `json:"path,omitempty" url:"path,omitempty"`
	Contents  string `json:"contents,omitempty" url:"contents,omitempty"`
	Overwrite bool   `json:"overwrite,omitempty" url:"overwrite,omitempty"`
}