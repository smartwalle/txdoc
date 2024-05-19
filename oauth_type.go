package txdoc

import "time"

// Token 获取 Access Token 返回数据 https://docs.qq.com/open/document/app/oauth2/access_token.html
type Token struct {
	Error
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	ExpiresTime  int64  `json:"-"`
	RefreshToken string `json:"refresh_token"`
	UserId       string `json:"user_id"`
	Scope        string `json:"scope"`
}

func (token *Token) Expired() bool {
	return token.ExpiresTime <= time.Now().Unix()
}

// GetUserResponse 获取用户信息返回数据 https://docs.qq.com/open/document/app/oauth2/user_info.html
type GetUserResponse struct {
	Error
	User
}

type User struct {
	OpenId  string `json:"openID"`
	Nick    string `json:"nick"`
	Avatar  string `json:"avatar"`
	Source  string `json:"source"`
	UnionId string `json:"unionID"`
}
