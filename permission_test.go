package txdoc_test

import (
	"context"
	"testing"
	"txdoc"
)

func TestAPI_GetCollaborators(t *testing.T) {
	var rsp, err = api.GetCollaborators(context.Background(), "sss")
	if err != nil {
		t.Fatal(err)
	}
	if rsp.IsFailure() {
		t.Fatal(rsp.Error)
	}
	t.Log(rsp)
}

func TestAPI_AddCollaborators(t *testing.T) {
	var param = txdoc.AddCollaboratorsParam{}
	param.Collaborators = append(
		param.Collaborators,
		txdoc.Collaborator{OpenId: "xxx", Type: txdoc.CollaboratorTypeUser, Role: txdoc.RoleTypeWriter},
	)

	var rsp, err = api.AddCollaborators(context.Background(), "sss", param)
	if err != nil {
		t.Fatal(err)
	}
	if rsp.IsFailure() {
		t.Fatal(rsp.Error)
	}
	t.Log(rsp)
}

func TestAPI_DeleteCollaborators(t *testing.T) {
	var rsp, err = api.DeleteCollaborators(context.Background(), "sss", "ddd")
	if err != nil {
		t.Fatal(err)
	}
	if rsp.IsFailure() {
		t.Fatal(rsp.Error)
	}
	t.Log(rsp)
}
