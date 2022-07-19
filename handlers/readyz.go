package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Readiness check to check if DB is running
func Readyz(started time.Time) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		duration := time.Since(started)
		if duration.Seconds() < 5 {
			log.Printf("readyz status: %v\n", http.StatusInternalServerError)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("error: %v", duration.Seconds()))) //nolint:errcheck
		} else {
			log.Printf("readyz status: %v\n", http.StatusOK)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("ok")) //nolint:errcheck
		}
	}
}
