package model_test

import (
    "testing"
    "time"

    model "github.com/rafaelbrunotech/general-server-go/internal/common/domain/model"
)

func TestNewApiRequest_AssignsFields(t *testing.T) {
    ts := time.Now()
    authToken := "token-123"
    r := model.NewApiRequest(authToken, ts)
    if r == nil {
        t.Fatalf("expected request instance, got nil")
    }
    if r.AuthToken != authToken {
        t.Fatalf("expected auth token set, got %q", r.AuthToken)
    }
    if r.Timestamp != ts {
        t.Fatalf("expected timestamp to be preserved")
    }
}