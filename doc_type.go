package txdoc

type DocumentType string

const (
	DocumentTypeDocument   DocumentType = "Document"   // Doc 文档
	DocumentTypeMainStory  DocumentType = "MainStory"  // 文档主节点，文档主体内容
	DocumentTypeSection    DocumentType = "Section"    // 节
	DocumentTypeParagraph  DocumentType = "Paragraph"  // 段落
	DocumentTypeTable      DocumentType = "Table"      // 表格
	DocumentTypeTableRow   DocumentType = "TableRow"   // 表格行
	DocumentTypeTableCell  DocumentType = "TableCell"  // 表格单元格
	DocumentTypeText       DocumentType = "Text"       // 一个具有相同属性集合的文本容器
	DocumentTypeDrawing    DocumentType = "Drawing"    // 图形化对象 , 例如图表、图片等
	DocumentTypeFootnote   DocumentType = "Footnote"   // 脚注
	DocumentTypeEndnote    DocumentType = "Endnote"    // 尾注
	DocumentTypeComment    DocumentType = "Comment"    // 批注
	DocumentTypeCommentRef DocumentType = "CommentRef" // 批注引用
	DocumentTypeTextBox    DocumentType = "TextBox"    // 文本框
	DocumentTypeFooter     DocumentType = "Footer"     // 页脚
	DocumentTypeHeader     DocumentType = "Header"     // 页眉
)

// Document https://docs.qq.com/open/document/app/openapi/v3/doc/get/document.html#
type Document struct {
	Begin    int          `json:"begin"`
	End      int          `json:"end"`
	Type     DocumentType `json:"type"`
	Children []Document   `json:"children"`
	Text     string       `json:"text"`
	//Property *DocumentProperty `json:"property,omitempty"` // 暂时不实现
}

type DocumentProperty struct {
}

// GetDocumentResponse 获取 Doc 文档内容返回数据 https://docs.qq.com/open/document/app/openapi/v3/doc/get/get.html
type GetDocumentResponse struct {
	Error
	Document
}
