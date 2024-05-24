package txdoc

import (
	"context"
	"net/http"
)

const (
	kPathV3Doc = "/openapi/doc/v3"
)

// GetDocument 获取 Doc 文档内容 https://docs.qq.com/open/document/app/openapi/v3/doc/get/get.html
func (api API) GetDocument(ctx context.Context, fileId string) (*GetDocumentResponse, error) {
	var aux = struct {
		ErrorV3
		Document GetDocumentResponse `json:"document"`
	}{}
	if err := api.request(ctx, http.MethodGet, api.buildAPI(kPathV3Doc, fileId), nil, nil, &aux); err != nil {
		return nil, err
	}
	aux.Document.Error = aux.ErrorV3.Error()
	return &aux.Document, nil
}
