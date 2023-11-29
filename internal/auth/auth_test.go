package auth

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

// GetAPIKey -
func TestGetAPIKeyFail(t *testing.T) {
	key, apiErr := GetAPIKey(http.Header{})
	want, wantErr := "", ErrNoAuthHeaderIncluded
	if key != want || apiErr != wantErr {
		t.Errorf("TestGetAPIKey() = (%v, %v), want (%v, %v)", key, apiErr, want, wantErr)
	}
	if !errors.Is(apiErr, wantErr) {
		t.Fatalf("expected: %v, apiErr: %v", wantErr, apiErr)
	}
}

// GetAPIKey -
func TestGetAPIKeySuccess(t *testing.T) {
	header := http.Header{}
	header.Add("Authorization", "ApiKey 123")
	key, apiErr := GetAPIKey(header)
	want := "123"

	if errors.Is(apiErr, ErrNoAuthHeaderIncluded) {
		t.Fatalf("expected: %v, key: %v", nil, apiErr)
	}
	if !reflect.DeepEqual(want, key) {
        t.Fatalf("expected: %v, key: %v", want, key)
    }
}