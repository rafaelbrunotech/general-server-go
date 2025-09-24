package model_test

import (
    "testing"

    model "github.com/rafaelbrunotech/general-server-go/internal/common/domain/model"
)

func TestNewSuccessApiResponse_SetsDataAndEnv(t *testing.T) {
    t.Setenv("ENV", "test")
    data := map[string]string{"ok": "yes"}
    r := model.NewSuccessApiResponse[map[string]string, map[string]string](&data, 200)
    if r.Status != 200 {
        t.Fatalf("expected status 200, got %d", r.Status)
    }
    if r.Error != nil {
        t.Fatalf("expected nil error, got %+v", r.Error)
    }
    if r.Metadata.Environment != "test" {
        t.Fatalf("expected environment from ENV, got %q", r.Metadata.Environment)
    }
    if (*r.Data)["ok"] != "yes" {
        t.Fatalf("expected data preserved, got %+v", r.Data)
    }
}

func TestNewErrorApiResponse_SetsErrorAndEnv(t *testing.T) {
    t.Setenv("ENV", "test")
    details := map[string]string{"field": "email"}
    r := model.NewErrorApiResponse[map[string]string, map[string]string](details, "invalid", 400)
    if r.Status != 400 {
        t.Fatalf("expected status 400, got %d", r.Status)
    }
    if r.Error == nil {
        t.Fatalf("expected error not nil")
    }
    if r.Error.Message != "invalid" {
        t.Fatalf("expected error message preserved")
    }
    if r.Error.Details["field"] != "email" {
        t.Fatalf("expected error details preserved")
    }
    if r.Metadata.Environment != "test" {
        t.Fatalf("expected environment from ENV, got %q", r.Metadata.Environment)
    }
    if r.Data != nil {
        t.Fatalf("expected data to be nil on error response")
    }
}