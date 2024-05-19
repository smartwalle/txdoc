package txdoc

import (
	"context"
	"io"
	"net/http"
	"net/url"
)

const (
	kHeaderAccessToken = "Access-Token"
	kHeaderClientId    = "client-Id"
	kHeaderOpenId      = "Open-Id"
)

type Document struct {
	client      *Client
	openId      string
	accessToken string
}

func (c *Client) Document(openId string, accessToken string) Document {
	var document = Document{}
	document.client = c
	document.openId = openId
	document.accessToken = accessToken
	return document
}

func (doc Document) request(ctx context.Context, method, url string, values url.Values, body io.Reader, result interface{}) error {
	var header = http.Header{}
	header.Set(kHeaderClientId, doc.client.clientId)
	header.Set(kHeaderOpenId, doc.openId)
	header.Set(kHeaderAccessToken, doc.accessToken)
	return doc.client.request(ctx, method, url, header, values, body, result)
}

func (doc Document) buildAPI(paths ...string) string {
	return doc.client.buildAPI(paths...)
}
