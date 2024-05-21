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
func (api API) GetCollaborators(ctx context.Context, fileId string) (*GetCollaboratorsResponse, error) {
	var aux = struct {
		Error
		Data GetCollaboratorsResponse `json:"data"`
	}{}
	if err := api.request(ctx, http.MethodGet, api.buildAPI(kPathV2Files, fileId, kPathCollaborators), nil, nil, &aux); err != nil {
		return nil, err
	}
	aux.Data.Error = aux.Error
	return &aux.Data, nil
}

// AddCollaborators 添加协作成员 https://docs.qq.com/open/document/app/openapi/v2/file/files/collaborators/add.html
func (api API) AddCollaborators(ctx context.Context, fileId string, param AddCollaboratorsParam) (*AddCollaboratorsResponse, error) {
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	var aux = struct {
		Error
		Data AddCollaboratorsResponse `json:"data"`
	}{}
	if err = api.request(ctx, http.MethodPatch, api.buildAPI(kPathV2Files, fileId, kPathCollaborators), nil, bytes.NewReader(data), &aux); err != nil {
		return nil, err
	}
	aux.Data.Error = aux.Error
	return &aux.Data, nil
}

// DeleteCollaborators 移除协作成员 https://docs.qq.com/open/document/app/openapi/v2/file/files/collaborators/delete.html
func (api API) DeleteCollaborators(ctx context.Context, fileId, openId string) (*DeleteCollaboratorsResponse, error) {
	var values = url.Values{}
	values.Set("type", string(CollaboratorTypeUser))
	values.Set("id", openId)

	var aux = struct {
		Error
		Data DeleteCollaboratorsResponse `json:"data"`
	}{}
	if err := api.request(ctx, http.MethodDelete, api.buildAPI(kPathV2Files, fileId, kPathCollaborators), values, nil, &aux); err != nil {
		return nil, err
	}
	aux.Data.Error = aux.Error
	return &aux.Data, nil
}
