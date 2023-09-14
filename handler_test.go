package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandler(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	require.NoError(t, err)

	rec := httptest.NewRecorder()

	Handle(rec, req)
	require.Equal(t, http.StatusOK, rec.Code)
	require.Equal(t, `{"status": "ok"}`, rec.Body.String())
}
