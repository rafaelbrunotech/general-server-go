package tokenizer_test

import (
    "testing"

    model "github.com/rafaelbrunotech/general-server-go/internal/common/domain/model"
    tokenizer "github.com/rafaelbrunotech/general-server-go/internal/common/infrastructure/service/tokenizer"
)

func TestTokenizer_GenerateAndVerifyTokens(t *testing.T) {
    t.Setenv("JWT_SECRET", "secret")
    tok, err := tokenizer.New()
    if err != nil { t.Fatalf("unexpected error: %v", err) }

    td := model.NewTokenData("id-1", "user@example.com")
    access, err := tok.GenerateAccessToken(*td)
    if err != nil || access == "" { t.Fatalf("failed to generate access: %v", err) }
    refresh, err := tok.GenerateRefreshToken(*td)
    if err != nil || refresh == "" { t.Fatalf("failed to generate refresh: %v", err) }

    if err := tok.VerifyToken(access); err != nil {
        t.Fatalf("verify access failed: %v", err)
    }
    if err := tok.VerifyToken(refresh); err != nil {
        t.Fatalf("verify refresh failed: %v", err)
    }

    decoded, err := tok.DecodeToken(access)
    if err != nil { t.Fatalf("decode failed: %v", err) }
    if decoded.UserId != "id-1" || decoded.UserEmail != "user@example.com" {
        t.Fatalf("decoded data mismatch: %+v", decoded)
    }
}

func TestTokenizer_VerifyToken_Invalid(t *testing.T) {
    t.Setenv("JWT_SECRET", "secret")
    tok, _ := tokenizer.New()
    if err := tok.VerifyToken("invalid.token.here"); err == nil {
        t.Fatalf("expected verification error for invalid token")
    }
}
