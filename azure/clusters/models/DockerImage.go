// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package models

type DockerImage struct {
	Url 				string						`json:"url,omitempty" url:"url,omitempty"`
	BasicAuth			DockerBasicAuth				`json:"basic_auth,omitempty" url:"basic_auth,omitempty"`
}