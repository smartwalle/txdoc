package txdoc

import (
	"context"
	"io"
	"net/http"
	"net/url"
)

const (
	kHeaderAccessToken = "Access-Token"
	kHeaderClientId    = "Client-Id"
	kHeaderOpenId      = "Open-Id"
)

type API struct {
	client      *Client
	openId      string
	accessToken string
}

func (c *Client) API(openId string, accessToken string) API {
	var document = API{}
	document.client = c
	document.openId = openId
	document.accessToken = accessToken
	return document
}

func (api API) request(ctx context.Context, method, url string, values url.Values, body io.Reader, result interface{}) error {
	var header = http.Header{}
	header.Set(kHeaderClientId, api.client.clientId)
	header.Set(kHeaderOpenId, api.openId)
	header.Set(kHeaderAccessToken, api.accessToken)
	return api.client.request(ctx, method, url, header, values, body, result)
}

func (api API) buildAPI(paths ...string) string {
	return api.client.buildAPI(paths...)
}
