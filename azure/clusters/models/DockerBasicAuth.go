// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package models

type DockerBasicAuth struct {
	Username			string					`json:"username,omitempty" url:"username,omitempty"`
	Password			string					`json:"password,omitempty" url:"password,omitempty"`
}