package txdoc_test

import (
	"context"
	"testing"
)

func TestAPI_GetSpreadsheet(t *testing.T) {
	var rsp, err = api.GetSpreadsheet(context.Background(), "fileid")
	if err != nil {
		t.Fatal(err)
	}
	if rsp.IsFailure() {
		t.Fatal(rsp.Error)
	}
	t.Log(rsp)
}
