package models

type ParamPair struct {
	Key   string `json:"key,omitempty" url:"key,omitempty"`
	Value string `json:"value,omitempty" url:"value,omitempty"`
}
