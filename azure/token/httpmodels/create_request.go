package httpmodels

type CreateReq struct {
	LifetimeSeconds int64  `json:"lifetime_seconds,omitempty" url:"lifetime_seconds,omitempty"`
	Comment         string `json:"comment,omitempty" url:"comment,omitempty"`
}
