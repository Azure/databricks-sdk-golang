package httpmodels

type MoveReq struct {
	SourcePath      string `json:"source_path,omitempty" url:"source_path,omitempty"`
	DestinationPath string `json:"destination_path,omitempty" url:"destination_path,omitempty"`
}