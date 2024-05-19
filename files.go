package txdoc

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
)

const (
	kPathFiles        = "/openapi/drive/v2/files"
	kPathFileMetadata = "/metadata"
)

func (doc Document) CreateFile(ctx context.Context, param CreateFileParam) (*CreateFileResponse, error) {
	var values = url.Values{}
	values.Set("title", param.Title)
	values.Set("type", string(param.FileType))
	values.Set("templateID", param.TemplateId)
	values.Set("templateVersion", param.TemplateVersion)
	values.Set("folderID", param.FolderId)
	values.Set("ext", param.Ext)

	var aux = struct {
		Error
		Data CreateFileResponse `json:"data"`
	}{}
	if err := doc.request(ctx, http.MethodPost, doc.buildAPI(kPathFiles), values, nil, &aux); err != nil {
		return nil, err
	}
	aux.Data.Error = aux.Error
	return &aux.Data, nil
}

// GetFile 查询文档 https://docs.qq.com/open/document/app/openapi/v2/file/files/metadata.html
func (doc Document) GetFile(ctx context.Context, fileId string) (*GetFileResponse, error) {
	var aux = struct {
		Error
		Data GetFileResponse `json:"data"`
	}{}
	if err := doc.request(ctx, http.MethodGet, doc.buildAPI(kPathFiles, fileId, kPathFileMetadata), nil, nil, &aux); err != nil {
		return nil, err
	}
	aux.Data.Error = aux.Error
	return &aux.Data, nil
}

// RenameFile 重命名文档 https://docs.qq.com/open/document/app/openapi/v2/file/files/update.html
func (doc Document) RenameFile(ctx context.Context, fileId, title string) (result *RenameFileResponse, err error) {
	var values = url.Values{}
	values.Set("title", title)

	if err = doc.request(ctx, http.MethodPatch, doc.buildAPI(kPathFiles, fileId), values, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteFile 删除文档 https://docs.qq.com/open/document/app/openapi/v2/file/files/delete.html
func (doc Document) DeleteFile(ctx context.Context, fileId string, sourceType SourceType, recoverable int) (result *DeleteFileResponse, err error) {
	var values = url.Values{}
	values.Set("type", string(sourceType))
	values.Set("recoverable", strconv.Itoa(recoverable))

	if err = doc.request(ctx, http.MethodDelete, doc.buildAPI(kPathFiles, fileId), values, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}
