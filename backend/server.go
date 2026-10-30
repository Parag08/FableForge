package main

import (
	"crypto/subtle"
	"encoding/json"
	"net/http"
	"time"

	"fableforge/internal/config"

	"github.com/golang-jwt/jwt/v5"
)

// Credentials is the JSON body accepted by POST /login.
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Claims are the JWT claims issued on successful login.
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// Server wires HTTP handlers to configuration.
type Server struct {
	cfg *config.Config
	now func() time.Time
}

// NewServer builds a Server from validated configuration.
func NewServer(cfg *config.Config) *Server {
	return &Server{cfg: cfg, now: time.Now}
}

// Routes returns the HTTP handler with all routes and CORS applied.
func (s *Server) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", s.handleHealth)
	mux.HandleFunc("/login", s.handleLogin)
	return s.cors(mux)
}

func (s *Server) cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Origin") == s.cfg.AllowedOrigin {
			h := w.Header()
			h.Set("Access-Control-Allow-Origin", s.cfg.AllowedOrigin)
			h.Set("Access-Control-Allow-Credentials", "true")
			h.Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			h.Set("Access-Control-Allow-Headers", "Content-Type")
			h.Add("Vary", "Origin")
		}
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		w.Header().Set("Allow", "GET, HEAD")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func (s *Server) handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var creds Credentials
	dec := json.NewDecoder(http.MaxBytesReader(w, r.Body, 1<<16))
	if err := dec.Decode(&creds); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	// Demo-only authentication against env-configured credentials.
	userOK := subtle.ConstantTimeCompare([]byte(creds.Username), []byte(s.cfg.DemoUsername)) == 1
	passOK := subtle.ConstantTimeCompare([]byte(creds.Password), []byte(s.cfg.DemoPassword)) == 1
	if !userOK || !passOK {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	expiresAt := s.now().Add(s.cfg.TokenTTL)
	claims := &Claims{
		Username: creds.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(s.now()),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(s.cfg.JWTSecret)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		Expires:  expiresAt,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
	w.WriteHeader(http.StatusOK)
}
