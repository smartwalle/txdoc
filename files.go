package txdoc

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
)

const (
	kPathV2Files      = "/openapi/drive/v2/files"
	kPathFileMetadata = "/metadata"
)

// CreateFile 新建文档 https://docs.qq.com/open/document/app/openapi/v2/file/files/create.html
func (api API) CreateFile(ctx context.Context, param CreateFileParam) (*CreateFileResponse, error) {
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
	if err := api.request(ctx, http.MethodPost, api.buildAPI(kPathV2Files), values, nil, &aux); err != nil {
		return nil, err
	}
	aux.Data.Error = aux.Error
	return &aux.Data, nil
}

// GetFile 查询文档 https://docs.qq.com/open/document/app/openapi/v2/file/files/metadata.html
func (api API) GetFile(ctx context.Context, fileId string) (*GetFileResponse, error) {
	var aux = struct {
		Error
		Data GetFileResponse `json:"data"`
	}{}
	if err := api.request(ctx, http.MethodGet, api.buildAPI(kPathV2Files, fileId, kPathFileMetadata), nil, nil, &aux); err != nil {
		return nil, err
	}
	aux.Data.Error = aux.Error
	return &aux.Data, nil
}

// RenameFile 重命名文档 https://docs.qq.com/open/document/app/openapi/v2/file/files/update.html
func (api API) RenameFile(ctx context.Context, fileId, title string) (result *RenameFileResponse, err error) {
	var values = url.Values{}
	values.Set("title", title)

	if err = api.request(ctx, http.MethodPatch, api.buildAPI(kPathV2Files, fileId), values, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteFile 删除文档 https://docs.qq.com/open/document/app/openapi/v2/file/files/delete.html
func (api API) DeleteFile(ctx context.Context, fileId string, sourceType SourceType, recoverable int) (result *DeleteFileResponse, err error) {
	var values = url.Values{}
	values.Set("type", string(sourceType))
	values.Set("recoverable", strconv.Itoa(recoverable))

	if err = api.request(ctx, http.MethodDelete, api.buildAPI(kPathV2Files, fileId), values, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}
