package txdoc_test

import (
	"context"
	"testing"
	"txdoc"
)

func TestAPI_GetCollaborators(t *testing.T) {
	var rsp, err = api.GetCollaborators(context.Background(), "300000000$FmmRKDlyqIlE")
	if err != nil {
		t.Fatal(err)
	}
	if rsp.IsFailure() {
		t.Fatal(rsp.Error)
	}
	t.Logf("%+v \n", rsp)
}

func TestAPI_AddCollaborators(t *testing.T) {
	var param = txdoc.AddCollaboratorsParam{}
	param.Collaborators = append(
		param.Collaborators,
		txdoc.Collaborator{OpenId: "84ce2862c86e4b3d8fd310cab984f8ba", Type: txdoc.CollaboratorTypeUser, Role: txdoc.RoleTypeWriter},
	)

	var rsp, err = api.AddCollaborators(context.Background(), "300000000$FteYXwBZzylK", param)
	if err != nil {
		t.Fatal(err)
	}
	if rsp.IsFailure() {
		t.Fatal(rsp.Error)
	}
	t.Logf("%+v \n", rsp)
}

func TestAPI_DeleteCollaborators(t *testing.T) {
	var rsp, err = api.DeleteCollaborators(context.Background(), "300000000$FmmRKDlyqIlE", "c1a19769e75a49158a381933b3aee84b")
	if err != nil {
		t.Fatal(err)
	}
	if rsp.IsFailure() {
		t.Fatal(rsp.Error)
	}
	t.Logf("%+v \n", rsp)
}
