package models

type NotebookOutput struct {
	Result    string `json:"result,omitempty" url:"result,omitempty"`
	Truncated bool   `json:"truncated,omitempty" url:"truncated,omitempty"`
}
