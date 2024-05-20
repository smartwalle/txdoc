package txdoc_test

import (
	"context"
	"testing"
)

func TestDocument_AsyncExport(t *testing.T) {
	t.Log(document.AsyncExport(context.Background(), "sss", ""))
}

func TestDocument_GetExportProgress(t *testing.T) {
	t.Log(document.GetExportProgress(context.Background(), "sss", "sss"))
}
