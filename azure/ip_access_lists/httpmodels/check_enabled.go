package httpmodels

type CheckEnabledReq struct {
	Key string `url:"key"`
}

type CheckEnabledResp struct {
	EnableIpAccessLists bool `json:"enableIpAccessLists,omitempty" url:"enableIpAccessLists,omitempty"`
}
