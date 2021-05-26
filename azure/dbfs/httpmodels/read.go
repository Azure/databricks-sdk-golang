package httpmodels

type ReadReq struct {
	Path   string `json:"path,omitempty" url:"path,omitempty"`
	Offset int64  `json:"offset,omitempty" url:"offset,omitempty"`
	Length int64  `json:"length,omitempty" url:"length,omitempty"`
}

type ReadResp struct {
	BytesRead int64  `json:"bytes_read,omitempty" url:"bytes_read,omitempty"`
	Data      string `json:"data,omitempty" url:"data,omitempty"`
}
