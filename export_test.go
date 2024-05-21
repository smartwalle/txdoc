package txdoc_test

import (
	"context"
	"testing"
)

func TestAPI_AsyncExport(t *testing.T) {
	var rsp, err = api.AsyncExport(context.Background(), "sss", "")
	if err != nil {
		t.Fatal(err)
	}
	if rsp.IsFailure() {
		t.Fatal(rsp.Error)
	}
	t.Log(rsp)
}

func TestAPI_GetExportProgress(t *testing.T) {
	var rsp, err = api.GetExportProgress(context.Background(), "sss", "sss")
	if err != nil {
		t.Fatal(err)
	}
	if rsp.IsFailure() {
		t.Fatal(rsp.Error)
	}
	t.Log(rsp)
}
