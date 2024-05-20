package txdoc

// ExportType 可导出的文档类型
type ExportType string

const (
	ExportTypeDoc   ExportType = "doc"
	ExportTypeSheet ExportType = "sheet"
	ExportTypeSlide ExportType = "slide"
	ExportTypePdf   ExportType = "pdf"
)

// AsyncExportResponse 导出文档返回数据 https://docs.qq.com/open/document/app/openapi/v2/file/export/async_export.html
type AsyncExportResponse struct {
	Error
	OperationId string `json:"operationID"`
}

// GetExportProgressResponse 导出进度查询返回数据 https://docs.qq.com/open/document/app/openapi/v2/file/export/export_progress.html
type GetExportProgressResponse struct {
	Error
	URL      string `json:"url"`
	Progress int    `json:"progress"`
}
