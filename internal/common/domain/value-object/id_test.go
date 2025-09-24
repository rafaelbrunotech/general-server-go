package valueobject_test

import (
    "testing"

    valueobject "github.com/rafaelbrunotech/general-server-go/internal/common/domain/value-object"
)

func TestNewId_GeneratesUUID(t *testing.T) {
    id := valueobject.NewId()
    if id == nil {
        t.Fatalf("expected id instance, got nil")
    }
    if id.Value() == "" {
        t.Fatalf("expected non-empty id value")
    }
}

func TestNewValue_ReturnsProvidedValue(t *testing.T) {
    const provided = "test-id-123"
    id := valueobject.NewValue(provided)
    if id.Value() != provided {
        t.Fatalf("expected %q, got %q", provided, id.Value())
    }
}