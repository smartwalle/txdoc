package txdoc_test

import (
	"context"
	"testing"
)

func TestAPI_GetDocument(t *testing.T) {
	var rsp, err = api.GetDocument(context.Background(), "fileid")
	if err != nil {
		t.Fatal(err)
	}
	if rsp.IsFailure() {
		t.Fatal(rsp.Error)
	}
	t.Log(rsp)
}
