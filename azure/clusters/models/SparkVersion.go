package models

type SparkVersion struct {
	Key  string `json:"key,omitempty" url:"key,omitempty"`
	Name string `json:"name,omitempty" url:"name,omitempty"`
}
