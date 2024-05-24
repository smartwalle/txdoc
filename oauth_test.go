package txdoc_test

import (
	"context"
	"testing"
)

func TestClient_GetAuthorizeURL(t *testing.T) {
	var rsp = client.GetAuthorizeURL("https://dev.ceo.qq.com", "test")
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
	t.Logf("%+v \n", rsp)
}

func TestClient_GetToken(t *testing.T) {
	var rsp, err = client.GetToken(context.Background(), "https://dev.ceo.qq.com", "WYTDTO5OOXKFVJ3ILY3P7A")
	if err != nil {
		t.Fatal(err)
	}
	if rsp.IsFailure() {
		t.Fatal(rsp.Error)
	}
	t.Logf("%+v \n", rsp)
}

func TestClient_RefreshToken(t *testing.T) {
	var rsp, err = client.RefreshToken(context.Background(), "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbHQiOiIwN2JiNTAyODkxYmQ0NzA3ODYyOTBlOGNlMGUxZTI4ZSIsInR5cCI6MiwidGV4cCI6MjU5MjAwMDAwMDAwMDAwMCwiZXhwIjoxNzQ4MDcxNTMzLjQxNDc4MiwiaWF0IjoxNzE2NTM1NTMzLjQxNDc4Miwic3ViIjoiODRjZTI4NjJjODZlNGIzZDhmZDMxMGNhYjk4NGY4YmEifQ.9v0K0IsBy40Gzqx03_DZwF14tDosT7irqrchwyMQOts")
	if err != nil {
		t.Fatal(err)
	}
	if rsp.IsFailure() {
		t.Fatal(rsp.Error)
	}
	t.Logf("%+v \n", rsp)
}

func TestClient_GetUserInfo(t *testing.T) {
	var rsp, err = client.GetUserInfo(context.Background(), "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbHQiOiIwN2JiNTAyODkxYmQ0NzA3ODYyOTBlOGNlMGUxZTI4ZSIsInR5cCI6MSwiZXhwIjoxNzE5MTI3NTc5LjI2NTQ4OTgsImlhdCI6MTcxNjUzNTU3OS4yNjU0ODk4LCJzdWIiOiI4NGNlMjg2MmM4NmU0YjNkOGZkMzEwY2FiOTg0ZjhiYSJ9.HI5Ys2z2rjcLLIIAo9xbqrGs1NDq_Nl6DjbnAnxR11c")
	if err != nil {
		t.Fatal(err)
	}
	if rsp.IsFailure() {
		t.Fatal(rsp.Error)
	}
	t.Logf("%+v \n", rsp)
}
