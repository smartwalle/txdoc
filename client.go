package txdoc

import (
	"context"
	"encoding/json"
	"github.com/smartwalle/ngx"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Option func(opts *Client)

func WithHTTPClient(client *http.Client) Option {
	return func(c *Client) {
		if client != nil {
			c.client = client
		}
	}
}

type Client struct {
	client       *http.Client
	host         string
	clientId     string
	clientSecret string
}

func New(clientId, clientSecret string, opts ...Option) *Client {
	var nClient = &Client{}
	nClient.client = http.DefaultClient
	nClient.host = kProductionGateway
	nClient.clientId = clientId
	nClient.clientSecret = clientSecret
	for _, opt := range opts {
		if opt != nil {
			opt(nClient)
		}
	}
	return nClient
}

func (c *Client) buildAPI(paths ...string) string {
	var path = c.host
	for _, p := range paths {
		p = strings.TrimSpace(p)
		if len(p) > 0 {
			if strings.HasSuffix(path, "/") {
				path = path + p
			} else {
				if strings.HasPrefix(p, "/") {
					path = path + p
				} else {
					path = path + "/" + p
				}
			}
		}
	}
	return path
}

func (c *Client) request(ctx context.Context, method, url string, header http.Header, values url.Values, body io.Reader, result interface{}) (err error) {
	var req = ngx.NewRequest(method, url, ngx.WithClient(c.client), ngx.WithHeader(header))
	if len(values) > 0 {
		req.SetForm(values)
	}
	if body != nil {
		req.SetBody(body)
		req.SetContentType(ngx.ContentTypeJSON)
	}

	rsp, err := req.Do(ctx)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	//fmt.Println(url, rsp.StatusCode)

	switch rsp.StatusCode {
	case http.StatusOK:
		return json.NewDecoder(rsp.Body).Decode(result)
	default:
		return Error{Code: Code(rsp.StatusCode), Message: http.StatusText(rsp.StatusCode)}
		//case http.StatusAccepted:
		//	return nil
		//case http.StatusUnauthorized:
		//	return &Error{Code: http.StatusUnauthorized, Message: http.StatusText(http.StatusUnauthorized)}
		//default:
		//	if len(data) == 0 {
		//		return &Error{Code: rsp.StatusCode, Message: http.StatusText(rsp.StatusCode)}
		//	}
		//
		//	var rErr *Error
		//	if err = json.Unmarshal(data, &rErr); err != nil {
		//		return err
		//	}
		//	return rErr
	}
	return nil
}
