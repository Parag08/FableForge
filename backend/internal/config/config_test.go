package config

import (
	"strings"
	"testing"
	"time"
)

func env(m map[string]string) func(string) (string, bool) {
	return func(k string) (string, bool) { v, ok := m[k]; return v, ok }
}

func valid() map[string]string {
	return map[string]string{
		KeyPort:          "8080",
		KeyJWTSecret:     strings.Repeat("x", 32),
		KeyTokenTTL:      "15m",
		KeyAllowedOrigin: "http://localhost:3000",
		KeyDemoUsername:  "demo",
		KeyDemoPassword:  "demo-pass",
	}
}

func TestLoadValid(t *testing.T) {
	cfg, err := LoadFrom(env(valid()))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Port != 8080 || cfg.TokenTTL != 15*time.Minute {
		t.Fatalf("unexpected config: %+v", cfg)
	}
}

func TestLoadMissingKeysFails(t *testing.T) {
	m := valid()
	delete(m, KeyJWTSecret)
	delete(m, KeyPort)
	_, err := LoadFrom(env(m))
	if err == nil {
		t.Fatal("expected error for missing keys")
	}
	for _, k := range []string{KeyJWTSecret, KeyPort} {
		if !strings.Contains(err.Error(), k) {
			t.Errorf("error %q does not mention %s", err, k)
		}
	}
}

func TestLoadInvalidValuesFail(t *testing.T) {
	m := valid()
	m[KeyTokenTTL] = "soon"
	m[KeyJWTSecret] = "short"
	if _, err := LoadFrom(env(m)); err == nil {
		t.Fatal("expected error for invalid values")
	}
}
