package model_test

import (
	"testing"

	model "github.com/rafaelbrunotech/general-server-go/internal/common/domain/model"
)

func TestNewTokenData_AssignsFields(t *testing.T) {
	td := model.NewTokenData("id-1", "user@example.com")
	if td.UserId != "id-1" || td.UserEmail != "user@example.com" {
		t.Fatalf("expected fields to be assigned, got: %+v", td)
	}
}
