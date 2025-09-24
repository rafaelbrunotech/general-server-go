package entity_test

import (
    "testing"
    "time"

    valueobject "github.com/rafaelbrunotech/general-server-go/internal/common/domain/value-object"
    entity "github.com/rafaelbrunotech/general-server-go/internal/packages/user/domain/entity"
)

func TestNewUser_Success(t *testing.T) {
    u, err := entity.NewUser(entity.UserInput{
        Email:    "USER@Example.com",
        Name:     "John",
        Password: "password123",
    })
    if err != nil { t.Fatalf("unexpected error: %v", err) }
    if u.Email.Value() != "user@example.com" { t.Fatalf("email normalized, got %q", u.Email.Value()) }
    if u.Name != "John" { t.Fatalf("name mismatch") }
    if u.Password == "password123" || u.Password == "" { t.Fatalf("password must be hashed and non-empty") }
    if u.CreatedAt.IsZero() || u.UpdatedAt.IsZero() { t.Fatalf("timestamps must be set") }
    if !u.IsPasswordCorrect("password123") { t.Fatalf("password verification failed") }
}

func TestNewUser_InvalidEmail(t *testing.T) {
    _, err := entity.NewUser(entity.UserInput{ Email: "invalid", Name: "John", Password: "pass" })
    if err == nil { t.Fatalf("expected error for invalid email") }
}

func TestUser_SetPassword_UpdatesHash(t *testing.T) {
    u, _ := entity.NewUser(entity.UserInput{ Email: "test@example.com", Name: "John", Password: "old" })
    oldHash := u.Password
    if err := u.SetPassword("newpass"); err != nil { t.Fatalf("set password error: %v", err) }
    if u.Password == oldHash { t.Fatalf("expected password hash to change") }
    if !u.IsPasswordCorrect("newpass") { t.Fatalf("new password should verify") }
}

func TestUser_SetEmail_And_SetName_UpdateTimestamps(t *testing.T) {
    u, _ := entity.NewUser(entity.UserInput{ Email: "t@example.com", Name: "John", Password: "p" })
    before := u.UpdatedAt
    time.Sleep(time.Millisecond) // ensure timestamp changes
    if err := u.SetEmail("new@example.com"); err != nil { t.Fatalf("set email error: %v", err) }
    if u.Email.Value() != "new@example.com" { t.Fatalf("email not updated") }
    if !u.UpdatedAt.After(before) { t.Fatalf("updatedAt should advance after SetEmail") }

    prev := u.UpdatedAt
    time.Sleep(time.Millisecond)
    u.SetName("Jane")
    if u.Name != "Jane" { t.Fatalf("name not updated") }
    if !u.UpdatedAt.After(prev) { t.Fatalf("updatedAt should advance after SetName") }
}

func TestUser_Restore_PopulatesFields(t *testing.T) {
    var u entity.User
    now := time.Now()
    input := entity.UserRestoreInput{
        Id:        "id-1",
        Email:     "user@example.com",
        Name:      "John",
        Password:  "hash",
        CreatedAt: now.Add(-time.Hour),
        UpdatedAt: now,
    }
    if err := u.Restore(input); err != nil { t.Fatalf("restore error: %v", err) }
    if u.Id.Value() != "id-1" { t.Fatalf("id mismatch") }
    if u.Email.Value() != "user@example.com" { t.Fatalf("email mismatch") }
    if u.Name != "John" || u.Password != "hash" { t.Fatalf("fields mismatch") }
    if !u.CreatedAt.Equal(input.CreatedAt) || !u.UpdatedAt.Equal(input.UpdatedAt) { t.Fatalf("timestamps mismatch") }
}

func TestUser_SetEmail_Invalid_ReturnsError(t *testing.T) {
    u, _ := entity.NewUser(entity.UserInput{ Email: "t@example.com", Name: "John", Password: "p" })
    err := u.SetEmail("invalid")
    if err == nil { t.Fatalf("expected error for invalid email") }
    // ensure previous value unchanged on error
    if u.Email.Value() != "t@example.com" { t.Fatalf("email should remain unchanged on error") }
}

func TestUser_Restore_InvalidEmail_Error(t *testing.T) {
    var u entity.User
    err := u.Restore(entity.UserRestoreInput{ Id: "id-1", Email: "invalid" })
    if err == nil { t.Fatalf("expected error") }
    _ = valueobject.EmailCannotBeEmpty // keep import if needed
}