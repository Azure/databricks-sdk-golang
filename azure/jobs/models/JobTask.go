package models

type JobTask struct {
	NotebookTask    NotebookTask    `json:"notebook_task,omitempty" url:"notebook_task,omitempty"`
	SparkJarTask    SparkJarTask    `json:"spark_jar_task,omitempty" url:"spark_jar_task,omitempty"`
	SparkPythonTask SparkPythonTask `json:"spark_python_task,omitempty" url:"spark_python_task,omitempty"`
	SparkSubmitTask SparkSubmitTask `json:"spark_submit_task,omitempty" url:"spark_submit_task,omitempty"`
}
