package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type InfoResponse struct {
	Time       string `json:"time"`
	Hostname   string `json:"hostname"`
	Message    string `json:"message"`
	DeployedOn string `json:"deployed_on"`
	AppEnv     string `json:"app_env"`
	AppName    string `json:"app_name"`
}

type HealthResponse struct {
	Status string `json:"status"`
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	response := InfoResponse{
		Time:       time.Now().Format("03:04:05PM on January 02, 2006"),
		Hostname:   hostname,
		Message:    "You are doing greattttttt, little human!!! <3",
		DeployedOn: "kubernetes",
		AppEnv:     "dev",
		AppName:    "go-app-1",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status: "up",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/api/v1/info", infoHandler)
	http.HandleFunc("/api/v1/healthz", healthHandler)

	addr := ":3000"
	log.Printf("Server starting on %s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
		os.Exit(1)
	}
}
