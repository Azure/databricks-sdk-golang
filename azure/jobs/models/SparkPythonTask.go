package models

type SparkPythonTask struct {
	PythonFile string    `json:"python_file,omitempty" url:"python_file,omitempty"`
	Parameters *[]string `json:"parameters,omitempty" url:"parameters,omitempty"`
}
