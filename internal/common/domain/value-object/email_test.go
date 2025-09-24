package valueobject_test

import (
    "testing"

    valueobject "github.com/rafaelbrunotech/general-server-go/internal/common/domain/value-object"
)

func TestNewEmail_Valid(t *testing.T) {
    e, err := valueobject.NewEmail("  USER@Example.COM  ")
    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }
    if e == nil {
        t.Fatalf("expected email instance, got nil")
    }
    if e.Value() != "user@example.com" {
        t.Fatalf("expected lowercased trimmed email, got %q", e.Value())
    }
}

func TestNewEmail_Empty(t *testing.T) {
    e, err := valueobject.NewEmail("   ")
    if err == nil {
        t.Fatalf("expected error, got nil")
    }
    if e != nil {
        t.Fatalf("expected nil email, got non-nil")
    }
    if err != valueobject.EmailCannotBeEmpty {
        t.Fatalf("expected EmailCannotBeEmpty, got %v", err)
    }
}

func TestNewEmail_InvalidFormat(t *testing.T) {
    tests := []string{
        "plainaddress",
        "@no-local-part.com",
        "user@",
        "user@com",
        "user@.com",
        "user@@example.com",
        "user example@example.com",
    }

    for _, tc := range tests {
        if e, err := valueobject.NewEmail(tc); err == nil || err != valueobject.EmailInvalidFormat || e != nil {
            t.Fatalf("expected invalid format error for %q", tc)
        }
    }
}