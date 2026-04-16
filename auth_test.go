package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestExtractTokenReadsGoTokenHeader(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/userinfo", nil)
	req.Header.Set("GoToken", "abc123")

	token := extractToken(req)
	if token != "abc123" {
		t.Fatalf("expected token from GoToken header, got %q", token)
	}
}

func TestExtractTokenReadsAuthorizationBearer(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/userinfo", nil)
	req.Header.Set("Authorization", "Bearer xyz987")

	token := extractToken(req)
	if token != "xyz987" {
		t.Fatalf("expected token from Authorization header, got %q", token)
	}
}

func TestLoginSuccessReturnsToken(t *testing.T) {
	body := `{"email":"user10@test.com","password":"test"}`
	req := httptest.NewRequest(http.MethodPost, "/api/login", strings.NewReader(body))
	res := httptest.NewRecorder()

	login(res, req)

	if res.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, res.Code)
	}

	var payload map[string]string
	if err := json.Unmarshal(res.Body.Bytes(), &payload); err != nil {
		t.Fatalf("expected json body, got error: %v", err)
	}
	if payload["token"] == "" {
		t.Fatal("expected token in response")
	}
}
