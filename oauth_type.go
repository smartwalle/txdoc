package txdoc

import "time"

// Token https://docs.qq.com/open/document/app/oauth2/access_token.html
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
