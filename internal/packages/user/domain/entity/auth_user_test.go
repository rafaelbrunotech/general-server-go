package entity_test

import (
    "testing"

    valueobject "github.com/rafaelbrunotech/general-server-go/internal/common/domain/value-object"
    entity "github.com/rafaelbrunotech/general-server-go/internal/packages/user/domain/entity"
)

func TestNewAuthUser_MapsFields(t *testing.T) {
    id := valueobject.NewValue("id-123")
    au := entity.NewAuthUser(entity.AuthUserInput{
        Id: id,
        Name: "John",
        AccessToken: "acc",
        RefreshToken: "ref",
    })
    if au.Id != "id-123" || au.Name != "John" || au.AccessToken != "acc" || au.RefreshToken != "ref" {
        t.Fatalf("unexpected mapping: %+v", au)
    }
}
