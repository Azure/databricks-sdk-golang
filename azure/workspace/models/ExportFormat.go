package models

type ExportFormat string

const (
	ExportFormatSource  = "SOURCE"
	ExportFormatHtml    = "HTML"
	ExportFormatJupyter = "JUPYTER"
	ExportFormatDbc     = "DBC"
)
