// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package models

type SparkJarTask struct {
	JarURI        string    `json:"jar_uri,omitempty" url:"jar_uri,omitempty"`
	MainClassName string    `json:"main_class_name,omitempty" url:"main_class_name,omitempty"`
	Parameters    *[]string `json:"parameters,omitempty" url:"parameters,omitempty"`
}
