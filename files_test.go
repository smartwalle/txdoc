package txdoc_test

import (
	"context"
	"testing"
	"txdoc"
)

func TestAPI_CreateFile(t *testing.T) {
	var param = txdoc.CreateFileParam{}
	param.Title = "testfile"
	param.FileType = txdoc.FileTypeSlide

	var rsp, err = api.CreateFile(context.Background(), param)
	if err != nil {
		t.Fatal(err)
	}
	if rsp.IsFailure() {
		t.Fatal(rsp.Error)
	}
	t.Logf("%+v \n", rsp)
}

func TestAPI_GetFile(t *testing.T) {
	var rsp, err = api.GetFile(context.Background(), "300000000$FmmRKDlyqIlE")
	if err != nil {
		t.Fatal(err)
	}
	if rsp.IsFailure() {
		t.Fatal(rsp.Error)
	}
	t.Logf("%+v \n", rsp)
}

func TestAPI_RenameFile(t *testing.T) {
	var rsp, err = api.RenameFile(context.Background(), "300000000$FNfywJFEQOCE", "newfile")
	if err != nil {
		t.Fatal(err)
	}
	if rsp.IsFailure() {
		t.Fatal(rsp.Error)
	}
	t.Logf("%+v \n", rsp)
}

func TestAPI_DeleteFile(t *testing.T) {
	var rsp, err = api.DeleteFile(context.Background(), "300000000$FNfywJFEQOCE", txdoc.SourceTypeOrigin, 0)
	if err != nil {
		t.Fatal(err)
	}
	if rsp.IsFailure() {
		t.Fatal(rsp.Error)
	}
	t.Logf("%+v \n", rsp)
}
