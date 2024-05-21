package txdoc

// GetSpreadsheetResponse 查询工作表信息返回数据 https://docs.qq.com/open/document/app/openapi/v3/sheet/get/get_sheet.html
type GetSpreadsheetResponse struct {
	Error
	SpreadsheetProperties
}

type SpreadsheetProperties struct {
	Properties  []SpreadsheetProperty `json:"properties"`
	RowTotal    int                   `json:"rowTotal"`
	ColumnTotal int                   `json:"columnTotal"`
}

type SpreadsheetProperty struct {
	SheetId     string `json:"sheetId"`
	Title       string `json:"title"`
	RowCount    int    `json:"rowCount"`
	ColumnCount int    `json:"columnCount"`
}
