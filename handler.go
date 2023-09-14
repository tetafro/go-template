package main

import "net/http"

// Handle handles http requests.
func Handle(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "ok"}`)) //nolint:errcheck,gosec
}
