// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package models

type RCranLibrary struct {
	Package string `json:"package,omitempty" url:"package,omitempty"`
	Repo    string `json:"repo,omitempty" url:"repo,omitempty"`
}
