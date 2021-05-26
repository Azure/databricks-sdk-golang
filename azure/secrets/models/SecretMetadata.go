package models

type SecretMetadata struct {
	Key                  string `json:"key,omitempty" url:"key,omitempty"`
	LastUpdatedTimestamp int64  `json:"last_updated_timestamp,omitempty" url:"last_updated_timestamp,omitempty"`
}
