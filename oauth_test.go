package txdoc_test

import (
	"context"
	"testing"
)

func TestClient_GetAppAccountToken(t *testing.T) {
	t.Log(client.GetAppAccountToken(context.Background()))
}

func TestClient_GetToken(t *testing.T) {
	t.Log(client.GetToken(context.Background(), "www", "ddd"))
}

func TestClient_RefreshToken(t *testing.T) {
	t.Log(client.RefreshToken(context.Background(), "sss"))
}
