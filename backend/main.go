package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"fableforge/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	addr := fmt.Sprintf(":%d", cfg.Port)
	srv := &http.Server{
		Addr:              addr,
		Handler:           NewServer(cfg).Routes(),
		ReadHeaderTimeout: 10 * time.Second,
	}

	log.Printf("FableForge backend listening on %s", addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
