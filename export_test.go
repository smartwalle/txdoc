package txdoc_test

import (
	"context"
	"testing"
	"txdoc"
)

func TestAPI_AsyncExport(t *testing.T) {
	var rsp, err = api.AsyncExport(context.Background(), "300000000$FteYXwBZzylK", txdoc.ExportTypeSlide)
	if err != nil {
		t.Fatal(err)
	}
	if rsp.IsFailure() {
		t.Fatal(rsp.Error)
	}
	t.Logf("%+v \n", rsp)
}

func TestAPI_GetExportProgress(t *testing.T) {
	var rsp, err = api.GetExportProgress(context.Background(), "300000000$FteYXwBZzylK", "b730d2ad396e43bcbd040ef6d0719609")
	if err != nil {
		t.Fatal(err)
	}
	if rsp.IsFailure() {
		t.Fatal(rsp.Error)
	}
	t.Logf("%+v \n", rsp)
}
