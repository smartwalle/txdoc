package txdoc

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	kPathCollaborators = "/collaborators"
)

// GetCollaborators 查询协作成员 https://docs.qq.com/open/document/app/openapi/v2/file/files/collaborators/get.html
func (doc Document) GetCollaborators(ctx context.Context, fileId string) (result *GetCollaboratorsResponse, err error) {
	var aux = struct {
		Error
		Data GetCollaboratorsResponse `json:"data"`
	}{}
	if err := doc.request(ctx, http.MethodGet, doc.buildAPI(kPathFiles, fileId, kPathCollaborators), nil, nil, &aux); err != nil {
		return nil, err
	}
	aux.Data.Error = aux.Error
	return &aux.Data, nil
}

// AddCollaborators 添加协作成员 https://docs.qq.com/open/document/app/openapi/v2/file/files/collaborators/add.html
func (doc Document) AddCollaborators(ctx context.Context, fileId string, param AddCollaboratorsParam) (result *AddCollaboratorsResponse, err error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	if err = doc.request(ctx, http.MethodPatch, doc.buildAPI(kPathFiles, fileId, kPathCollaborators), nil, bytes.NewReader(data), &result); err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteCollaborators 移除协作成员 https://docs.qq.com/open/document/app/openapi/v2/file/files/collaborators/delete.html
func (doc Document) DeleteCollaborators(ctx context.Context, fileId, openId string) (result *DeleteCollaboratorsResponse, err error) {
	var values = url.Values{}
	values.Set("type", string(CollaboratorTypeUser))
	values.Set("id", openId)

	if err = doc.request(ctx, http.MethodDelete, doc.buildAPI(kPathFiles, fileId, kPathCollaborators), values, nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}
