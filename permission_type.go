package txdoc

// RoleType 协作成员身份
type RoleType string

const (
	RoleTypeReader RoleType = "reader"
	RoleTypeWriter RoleType = "writer"
)

// CollaboratorType 协作类型
type CollaboratorType string

const (
	CollaboratorTypeUser CollaboratorType = "user"
)

type Collaborator struct {
	Type   CollaboratorType `json:"type"`
	Role   RoleType         `json:"role"`
	OpenId string           `json:"id"`
}

// GetCollaboratorsResponse 查询协作成员返回数据 https://docs.qq.com/open/document/app/openapi/v2/file/files/collaborators/get.html
type GetCollaboratorsResponse struct {
	Error
	Collaborators []Collaborator `json:"collaborators"`
}

// AddCollaboratorsParam 添加协作成员请求参数 https://docs.qq.com/open/document/app/openapi/v2/file/files/collaborators/add.html
type AddCollaboratorsParam struct {
	Collaborators []Collaborator `json:"collaborators"`
}

// AddCollaboratorsResponse 添加协作成员返回数据 https://docs.qq.com/open/document/app/openapi/v2/file/files/collaborators/add.html
type AddCollaboratorsResponse struct {
	Error
}

// DeleteCollaboratorsResponse 移除协作成员返回数据 https://docs.qq.com/open/document/app/openapi/v2/file/files/collaborators/delete.html
type DeleteCollaboratorsResponse struct {
	Error
}
