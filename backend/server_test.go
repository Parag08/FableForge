package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"fableforge/internal/config"

	"github.com/golang-jwt/jwt/v5"
)

func testConfig() *config.Config {
	return &config.Config{
		Port:          8080,
		JWTSecret:     []byte(strings.Repeat("s", 32)),
		TokenTTL:      15 * time.Minute,
		AllowedOrigin: "http://localhost:3000",
		DemoUsername:  "demo",
		DemoPassword:  "demo-pass",
	}
}

func do(t *testing.T, h http.Handler, method, path, body string) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	return rec
}

func TestHealthz(t *testing.T) {
	rec := do(t, NewServer(testConfig()).Routes(), http.MethodGet, "/healthz", "")
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), `"ok"`) {
		t.Fatalf("body = %q", rec.Body.String())
	}
}

func TestLoginSuccess(t *testing.T) {
	cfg := testConfig()
	rec := do(t, NewServer(cfg).Routes(), http.MethodPost, "/login", `{"username":"demo","password":"demo-pass"}`)
	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", rec.Code)
	}
	var tokenCookie *http.Cookie
	for _, c := range rec.Result().Cookies() {
		if c.Name == "token" {
			tokenCookie = c
		}
	}
	if tokenCookie == nil || !tokenCookie.HttpOnly {
		t.Fatalf("expected HttpOnly token cookie, got %+v", rec.Result().Cookies())
	}
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(tokenCookie.Value, claims, func(*jwt.Token) (any, error) { return cfg.JWTSecret, nil })
	if err != nil {
		t.Fatalf("token did not verify: %v", err)
	}
	if claims.Username != "demo" {
		t.Fatalf("username claim = %q", claims.Username)
	}
}

func TestLoginFailure(t *testing.T) {
	h := NewServer(testConfig()).Routes()
	cases := []struct {
		name, method, body string
		want               int
	}{
		{"wrong password", http.MethodPost, `{"username":"demo","password":"nope"}`, http.StatusUnauthorized},
		{"wrong user", http.MethodPost, `{"username":"x","password":"demo-pass"}`, http.StatusUnauthorized},
		{"malformed json", http.MethodPost, `{not json`, http.StatusBadRequest},
		{"wrong method", http.MethodGet, ``, http.StatusMethodNotAllowed},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			rec := do(t, h, tc.method, "/login", tc.body)
			if rec.Code != tc.want {
				t.Fatalf("status = %d, want %d", rec.Code, tc.want)
			}
			for _, c := range rec.Result().Cookies() {
				if c.Name == "token" {
					t.Fatal("token cookie set on failed login")
				}
			}
		})
	}
}
