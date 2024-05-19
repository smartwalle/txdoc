package txdoc

// FileType 文档类型
type FileType string

const (
	FileTypeDoc        FileType = "doc"
	FileTypeSheet      FileType = "sheet"
	FileTypeForm       FileType = "form"
	FileTypeSlide      FileType = "slide"
	FileTypeMind       FileType = "mind"
	FileTypeFlowchart  FileType = "flowchart"
	FileTypeAddon      FileType = "addon"
	FileTypeSmartSheet FileType = "smartsheet"
)

// SourceType 文件所属的列表类型
type SourceType string

const (
	SourceTypeOrigin SourceType = "origin"
	SourceTypeShared SourceType = "shared"
	SourceTypeRecent SourceType = "recent"
	SourceTypeTrash  SourceType = "trash"
)

// File 文件信息 https://docs.qq.com/open/document/app/openapi/v2/file/common/file_info.html
type File struct {
	ID                    string `json:"ID"`
	Title                 string `json:"title"`
	Type                  string `json:"type"`
	URL                   string `json:"url"`
	Status                string `json:"status"`
	IsCreator             bool   `json:"isCreator"`
	CreateTime            int64  `json:"createTime"`
	CreatorName           string `json:"creatorName"`
	IsOwner               bool   `json:"isOwner"`
	OwnerName             string `json:"ownerName"`
	LastModifyTime        int64  `json:"lastModifyTime"`
	LastModifyName        string `json:"lastModifyName"`
	FormCollectingStatus  string `json:"formCollectingStatus"`
	FormCollectingEndTime string `json:"formCollectingEndTime"`
}

// CreateFileParam 新建文档参数 https://docs.qq.com/open/document/app/openapi/v2/file/files/create.html
type CreateFileParam struct {
	Title           string   `json:"title"`
	FileType        FileType `json:"fileType"`
	TemplateId      string   `json:"templateID"`
	TemplateVersion string   `json:"templateVersion"`
	FolderId        string   `json:"folderID"`
	Ext             string   `json:"ext"`
}

// CreateFileResponse 新建文档返回数据 https://docs.qq.com/open/document/app/openapi/v2/file/files/create.html
type CreateFileResponse struct {
	Error
	File
}

// GetFileResponse 查询文档返回数据 https://docs.qq.com/open/document/app/openapi/v2/file/files/metadata.html
type GetFileResponse struct {
	Error
	File
}

// RenameFileResponse 重命名文档返回数据 https://docs.qq.com/open/document/app/openapi/v2/file/files/update.html
type RenameFileResponse struct {
	Error
}

// DeleteFileResponse 删除文档返回数据 https://docs.qq.com/open/document/app/openapi/v2/file/files/delete.html
type DeleteFileResponse struct {
	Error
}
