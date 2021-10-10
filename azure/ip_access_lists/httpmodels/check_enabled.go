package httpmodels

type CheckEnabledReq struct {
	Keys string `url:"keys"`
}

type CheckEnabledResp struct {
	EnableIpAccessLists string `json:"enableIpAccessLists,omitempty" url:"enableIpAccessLists,omitempty"`
}
