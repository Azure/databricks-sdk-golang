package httpmodels

type DeleteSecretScopeReq struct {
	Scope string `json:"scope,omitempty" url:"scope,omitempty"`
}
