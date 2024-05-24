package txdoc

import (
	"context"
	"net/http"
)

const (
	kPathV3Spreadsheet = "/openapi/spreadsheet/v3/files"
)

// GetSpreadsheet 查询工作表信息 https://docs.qq.com/open/document/app/openapi/v3/sheet/get/get_sheet.html
func (api API) GetSpreadsheet(ctx context.Context, fileId string) (*GetSpreadsheetResponse, error) {
	var aux = struct {
		ErrorV3
		GetSpreadsheetResponse
	}{}
	if err := api.request(ctx, http.MethodGet, api.buildAPI(kPathV3Spreadsheet, fileId), nil, nil, &aux); err != nil {
		return nil, err
	}
	aux.GetSpreadsheetResponse.Error = aux.ErrorV3.Error()
	return &aux.GetSpreadsheetResponse, nil
}
