package txdoc

import (
	"context"
	"net/http"
	"net/url"
	"time"
)

const (
	kPathOAuthAccountToken = "/oauth/v2/app-account-token"
	kPathOAuthToken        = "/oauth/v2/token"
)

const (
	kGrantTypeAuthorizationCode = "authorization_code"
	kGrantTypeRefreshToken      = "refresh_token"
)

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
