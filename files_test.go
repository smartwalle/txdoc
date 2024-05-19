package txdoc_test

import (
	"context"
	"testing"
	"txdoc"
)

func TestClient_CreateFile(t *testing.T) {
	var param = txdoc.CreateFileParam{}
	param.Title = "testfile"
	var rsp, err = document.CreateFile(context.Background(), param)
	if err != nil {
		t.Fatal(err)
	}
	if rsp.IsFailure() {
		t.Fatal(rsp.Error)
	}
	t.Log(rsp.Title)
}

func TestClient_GetFile(t *testing.T) {
	var rsp, err = document.GetFile(context.Background(), "fileid")
	if err != nil {
		t.Fatal(err)
	}
	if rsp.IsFailure() {
		t.Fatal(rsp.Error)
	}
	t.Log(rsp.Title)
}
