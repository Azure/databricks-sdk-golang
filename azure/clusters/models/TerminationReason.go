package models

type TerminationReason struct {
	Code       TerminationCode `json:"code,omitempty" url:"code,omitempty"`
	Type       TerminationType `json:"type,omitempty" url:"code,omitempty"`
	Parameters []ParameterPair `json:"parameters,omitempty" url:"parameters,omitempty"`
}
