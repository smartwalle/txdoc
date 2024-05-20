package txdoc

import (
	"context"
	"net/http"
	"net/url"
)

const (
	kPathAsyncExport    = "/async-export"
	kPathExportProgress = "/export-progress"
)

// AsyncExport 导出文档 https://docs.qq.com/open/document/app/openapi/v2/file/export/async_export.html
func (doc Document) AsyncExport(ctx context.Context, fileId string, exportType ExportType) (*AsyncExportResponse, error) {
	var values = url.Values{}
	values.Set("exportType", string(exportType))

	var aux = struct {
		Error
		Data AsyncExportResponse `json:"data"`
	}{}
	if err := doc.request(ctx, http.MethodPost, doc.buildAPI(kPathFiles, fileId, kPathAsyncExport), values, nil, &aux); err != nil {
		return nil, err
	}
	aux.Data.Error = aux.Error
	return &aux.Data, nil
}

// GetExportProgress 导出进度查询 https://docs.qq.com/open/document/app/openapi/v2/file/export/export_progress.html
func (doc Document) GetExportProgress(ctx context.Context, fileId, operationId string) (*GetExportProgressResponse, error) {
	var values = url.Values{}
	values.Set("operationID", operationId)

	var aux = struct {
		Error
		Data GetExportProgressResponse `json:"data"`
	}{}
	if err := doc.request(ctx, http.MethodGet, doc.buildAPI(kPathFiles, fileId, kPathExportProgress), values, nil, &aux); err != nil {
		return nil, err
	}
	aux.Data.Error = aux.Error
	return &aux.Data, nil
}
