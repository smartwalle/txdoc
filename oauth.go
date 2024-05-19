package txdoc

import (
	"context"
	"net/http"
	"net/url"
	"time"
)

const (
	kPathAuthorize         = "/oauth/v2/authorize"
	kPathOAuthAccountToken = "/oauth/v2/app-account-token"
	kPathOAuthToken        = "/oauth/v2/token"
	kPathUserInfo          = "/oauth/v2/userinfo"
)

const (
	kGrantTypeAuthorizationCode = "authorization_code"
	kGrantTypeRefreshToken      = "refresh_token"
)

// GetAuthorizeURL 获取前端发起授权URL https://docs.qq.com/open/document/app/oauth2/authorize.html
func (c *Client) GetAuthorizeURL(redirectURI, state string) *url.URL {
	var values = url.Values{}
	values.Set("client_id", c.clientId)
	values.Set("redirect_uri", redirectURI)
	values.Set("response_type", "code")
	values.Set("scope", "all")
	values.Set("state", state)
	var nURL, _ = url.Parse(c.buildAPI(kPathAuthorize) + "?" + values.Encode())
	return nURL
}

// GetAppAccountToken 获取应用级账号 Token https://docs.qq.com/open/document/app/oauth2/app_account_token.html
func (c *Client) GetAppAccountToken(ctx context.Context) (token *Token, err error) {
	var values = url.Values{}
	values.Set("client_id", c.clientId)
	values.Set("client_secret", c.clientSecret)
	if err = c.request(ctx, http.MethodGet, c.buildAPI(kPathOAuthAccountToken), nil, values, nil, &token); err != nil {
		return nil, err
	}
	if token.IsSuccess() {
		token.ExpiresTime = time.Now().Unix() + int64(token.ExpiresIn) - 30 // token 过期时间减少 30 秒
	}
	return token, nil
}

// GetToken 获取 Access Token https://docs.qq.com/open/document/app/oauth2/access_token.html
func (c *Client) GetToken(ctx context.Context, redirectURI, code string) (token *Token, err error) {
	var values = url.Values{}
	values.Set("client_id", c.clientId)
	values.Set("client_secret", c.clientSecret)
	values.Set("redirect_uri", redirectURI)
	values.Set("grant_type", kGrantTypeAuthorizationCode)
	values.Set("code", code)
	if err = c.request(ctx, http.MethodGet, c.buildAPI(kPathOAuthToken), nil, values, nil, &token); err != nil {
		return nil, err
	}
	if token.IsSuccess() {
		token.ExpiresTime = time.Now().Unix() + int64(token.ExpiresIn) - 30 // token 过期时间减少 30 秒
	}
	return token, nil
}

// RefreshToken 刷新 Token https://docs.qq.com/open/document/app/oauth2/refresh_token.html
func (c *Client) RefreshToken(ctx context.Context, refreshToken string) (token *Token, err error) {
	var values = url.Values{}
	values.Set("client_id", c.clientId)
	values.Set("client_secret", c.clientSecret)
	values.Set("grant_type", kGrantTypeRefreshToken)
	values.Set("refresh_token", refreshToken)
	if err = c.request(ctx, http.MethodGet, c.buildAPI(kPathOAuthToken), nil, values, nil, &token); err != nil {
		return nil, err
	}
	token.RefreshToken = refreshToken
	token.ExpiresTime = time.Now().Unix() + int64(token.ExpiresIn) - 30 // token 过期检测时间减少 30 秒
	return token, nil
}

// GetUserInfo 获取用户信息 https://docs.qq.com/open/document/app/oauth2/user_info.html
func (c *Client) GetUserInfo(ctx context.Context, accessToken string) (*GetUserResponse, error) {
	var values = url.Values{}
	values.Set("access_token", accessToken)

	var aux = struct {
		Error
		Data GetUserResponse `json:"data"`
	}{}
	if err := c.request(ctx, http.MethodGet, c.buildAPI(kPathUserInfo), nil, values, nil, &aux); err != nil {
		return nil, err
	}
	aux.Data.Error = aux.Error
	return &aux.Data, nil
}
