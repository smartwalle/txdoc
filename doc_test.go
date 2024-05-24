package txdoc_test

import (
	"context"
	"testing"
)

func TestAPI_GetDocument(t *testing.T) {
	var rsp, err = api.GetDocument(context.Background(), "300000000$FmmRKDlyqIlE")
	if err != nil {
		t.Fatal(err)
	}
	if rsp.IsFailure() {
		t.Fatal(rsp.Error)
	}
	t.Logf("%+v \n", rsp)
}
