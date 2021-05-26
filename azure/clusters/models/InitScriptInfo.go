package models

type InitScriptInfo struct {
	Dbfs DbfsStorageInfo `json:"dbfs,omitempty" url:"dbfs,omitempty"`
	Fs   FileStorageInfo `json:"fs,omitempty" url:"fs,omitempty"`
}
