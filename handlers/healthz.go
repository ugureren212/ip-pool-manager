package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Health check to check if server is running
func Healthz(started time.Time) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		duration := time.Since(started)
		if duration.Seconds() < 10 {
			log.Printf("healthz status: %v\n", http.StatusInternalServerError)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("error: %v", duration.Seconds()))) //nolint:errcheck
		} else {
			log.Printf("healthz status: %v\n", http.StatusOK)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("ok")) //nolint:errcheck
		}
	}
}
