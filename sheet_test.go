package txdoc_test

import (
	"context"
	"testing"
)

func TestAPI_GetSpreadsheet(t *testing.T) {
	var rsp, err = api.GetSpreadsheet(context.Background(), "300000000$FARfiOEEVieI")
	if err != nil {
		t.Fatal(err)
	}
	if rsp.IsFailure() {
		t.Fatal(rsp.Error)
	}
	t.Logf("%+v \n", rsp)
}
