package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// version is set at build time: -ldflags "-X main.version=<git sha>"
var version = "dev"

const name = "${{ values.name }}"

func main() {
	message := os.Getenv("APP_MESSAGE")
	if message == "" {
		message = "Hello from " + name
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, map[string]string{"service": name, "message": message})
	})
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, map[string]string{"status": "ok"})
	})
	mux.HandleFunc("GET /version", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, map[string]string{"version": version})
	})

	log.Printf("%s %s listening on :8080", name, version)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(v)
}
