package models

type ClusterLogConf struct {
	Dbfs DbfsStorageInfo `json:"dbfs,omitempty" url:"dbfs,omitempty"`
}
