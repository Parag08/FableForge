// Package config loads all runtime settings from environment variables.
// There are deliberately no in-code defaults: every key is required and
// Load fails fast, naming every missing or invalid key.
package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Environment variable names.
const (
	KeyPort          = "PORT"
	KeyJWTSecret     = "JWT_SECRET"
	KeyTokenTTL      = "TOKEN_TTL"
	KeyAllowedOrigin = "ALLOWED_ORIGIN"
	KeyDemoUsername  = "DEMO_USERNAME"
	KeyDemoPassword  = "DEMO_PASSWORD"
)

// Config holds the validated runtime settings.
type Config struct {
	Port          int
	JWTSecret     []byte
	TokenTTL      time.Duration
	AllowedOrigin string
	DemoUsername  string
	DemoPassword  string
}

// Load reads configuration from the process environment.
func Load() (*Config, error) {
	return LoadFrom(os.LookupEnv)
}

// LoadFrom reads configuration using the supplied lookup function
// (useful in tests).
func LoadFrom(lookup func(string) (string, bool)) (*Config, error) {
	var problems []string
	get := func(key string) string {
		v, ok := lookup(key)
		v = strings.TrimSpace(v)
		if !ok || v == "" {
			problems = append(problems, key+" is required")
		}
		return v
	}

	cfg := &Config{}

	if raw := get(KeyPort); raw != "" {
		p, err := strconv.Atoi(raw)
		if err != nil || p < 1 || p > 65535 {
			problems = append(problems, KeyPort+" must be an integer in 1-65535")
		}
		cfg.Port = p
	}
	if raw := get(KeyJWTSecret); raw != "" {
		if len(raw) < 32 {
			problems = append(problems, KeyJWTSecret+" must be at least 32 characters")
		}
		cfg.JWTSecret = []byte(raw)
	}
	if raw := get(KeyTokenTTL); raw != "" {
		d, err := time.ParseDuration(raw)
		if err != nil || d <= 0 {
			problems = append(problems, KeyTokenTTL+" must be a positive Go duration (e.g. 15m)")
		}
		cfg.TokenTTL = d
	}
	cfg.AllowedOrigin = get(KeyAllowedOrigin)
	cfg.DemoUsername = get(KeyDemoUsername)
	cfg.DemoPassword = get(KeyDemoPassword)

	if len(problems) > 0 {
		return nil, fmt.Errorf("invalid configuration: %s", strings.Join(problems, "; "))
	}
	return cfg, nil
}
