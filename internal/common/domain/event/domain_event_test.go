package event_test

import (
	"testing"
	"time"

	event "github.com/rafaelbrunotech/general-server-go/internal/common/domain/event"
)

func TestNewDomainEvent_AssignsDefaults(t *testing.T) {
	before := time.Now()
	eventName := "user_signed_up"
	e := event.NewDomainEvent(map[string]string{"a": "b"}, eventName)

	if e.Name != eventName {
		t.Fatalf("expected name set")
	}
	if e.Version != 1 {
		t.Fatalf("expected default version 1")
	}
	if e.Date.Before(before) {
		t.Fatalf("expected event date to be now-ish")
	}
	m, ok := e.Data.(map[string]string)
	if !ok || m["a"] != "b" {
		t.Fatalf("expected data preserved")
	}
}
