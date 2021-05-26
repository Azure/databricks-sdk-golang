package httpmodels

type PutSecretReq struct {
	StringValue string `json:"string_value,omitempty" url:"string_value,omitempty"`
	BytesValue  string `json:"bytes_value,omitempty" url:"bytes_value,omitempty"`
	Scope       string `json:"scope,omitempty" url:"scope,omitempty"`
	Key         string `json:"key,omitempty" url:"key,omitempty"`
}
