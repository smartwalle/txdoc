package txdoc_test

import (
	"context"
	"testing"
)

func TestClient_GetAuthorizeURL(t *testing.T) {
	var rsp = client.GetAuthorizeURL("http://www.baidu.com", "xxx")
	t.Log(rsp)
}

func TestClient_GetAppAccountToken(t *testing.T) {
	var rsp, err = client.GetAppAccountToken(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if rsp.IsFailure() {
		t.Fatal(rsp.Error)
	}
	t.Log(rsp)
}

func TestClient_GetToken(t *testing.T) {
	var rsp, err = client.GetToken(context.Background(), "www", "ddd")
	if err != nil {
		t.Fatal(err)
	}
	if rsp.IsFailure() {
		t.Fatal(rsp.Error)
	}
	t.Log(rsp)
}

func TestClient_RefreshToken(t *testing.T) {
	var rsp, err = client.RefreshToken(context.Background(), "sss")
	if err != nil {
		t.Fatal(err)
	}
	if rsp.IsFailure() {
		t.Fatal(rsp.Error)
	}
	t.Log(rsp)
}

func TestClient_GetUserInfo(t *testing.T) {
	var rsp, err = client.GetUserInfo(context.Background(), "sss")
	if err != nil {
		t.Fatal(err)
	}
	if rsp.IsFailure() {
		t.Fatal(rsp.Error)
	}
	t.Log(rsp)
}
