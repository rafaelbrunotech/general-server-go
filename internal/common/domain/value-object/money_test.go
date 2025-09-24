package valueobject_test

import (
    "testing"

    valueobject "github.com/rafaelbrunotech/general-server-go/internal/common/domain/value-object"
)

func TestNewMoney_DefaultsToZero(t *testing.T) {
    m, err := valueobject.NewMoney()
    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }
    if m == nil {
        t.Fatalf("expected money instance, got nil")
    }
    if m.Value() != 0 {
        t.Fatalf("expected default value 0, got %d", m.Value())
    }
}